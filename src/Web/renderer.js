var board = [];

var BLOCKED = 3;
var OPEN = 0;
var START = 1;
var TARGET = 2;
var PATH = 4;

var cellSize;

//init random problem by size
CreateBoard = function (size) {
    board = [];
    for (let y = 0; y < size; y++) {
        var row = [];
        for (let y = 0; y < size; y++) {
            //%30 chance to create closed block
            row.push(Math.random() < .3 ? BLOCKED : OPEN)
        }
        board.push(row);
    }

    //Pick random start point
    var rx = Math.floor(Math.random() * size);
    var ry = Math.floor(Math.random() * size);
    board[ry][rx] = START;

}

//Visualise  data
RenderBoard = function (m) {
    var c = document.getElementById("board");
    var ctx = c.getContext("2d");
    cellSize = c.width / m.length;
    for (let i = 0; i < m.length; i++) {
        var row = m[i];
        for (let j = 0; j < row.length; j++) {
            var val = row[j]
            ctx.lineWidth = 1;
            switch (val) {
                case OPEN:
                        case PATH:
                    ctx.fillStyle = "#ffffff";
                    break;
                case START:
                    ctx.fillStyle = "#0000ff";
                    break;
                case TARGET:
                    ctx.fillStyle = "#ff0000";
                    break;            
                case BLOCKED:
                    ctx.fillStyle = "#cccccc";
                    break;
                default:
                    break;
            }
            ctx.strokeRect(j * cellSize, i * cellSize, cellSize, cellSize);
            ctx.fillRect(j * cellSize, i * cellSize, cellSize, cellSize);
            if (val==PATH){
                //solution path
                ctx.lineWidth = 0;
                ctx.beginPath();
                ctx.arc(j * cellSize+cellSize/2, i * cellSize+cellSize/2, cellSize/4, 0, 2 * Math.PI);
                ctx.fillStyle = "#00ff00";
                ctx.fill();
                ctx.stroke();
            }
        }
    }
}

//set new destionation point for problem
BoardClick = function (event) {
    //revert canvas
    RenderBoard(board);
    
    //Deep copy problem. A faster way needed..
    var m = JSON.parse(JSON.stringify(board));

    var Destination = {};
    var size = m.length;

    //Map clicked cell
    Destination.x = Math.floor(event.offsetX / cellSize);
    Destination.y = Math.floor(event.offsetY / cellSize);
    
    //Open nodes only
    if (m[Destination.y][Destination.x] != OPEN) {
        return
    }

    m[Destination.y][Destination.x] = TARGET;
    //One dimension array to send  WASM app
    var arr = [m.length, m.length];
    for (let y = 0; y < m.length; y++) {
        for (let x = 0; x < m.length; x++) {
            arr.push(m[y][x]);
        }
    }

    //Actual A* run on WASM app
    var solution = Resolve(...arr);
    
    if (!solution.length) return;
    for (let i = 1; i < solution.length; i++) {
        var s = solution[i];
        m[Math.floor(s / size)][s % size] = PATH;
    } 
    RenderBoard(m);
}

InitBoard=function(){
    CreateBoard(20);
    RenderBoard(board);
}

document.addEventListener("DOMContentLoaded", function () {
    var c = document.getElementById('board');
    c.addEventListener('click', BoardClick, false);
    InitBoard();
});
