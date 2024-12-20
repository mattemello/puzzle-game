const go = new Go();
//wasm.CreateTheArena();

document.addEventListener('keydown', function(event) {

    if (event.key == "L" || event.key == "D" || (event.shiftKey && event.key == "ArrowRight")) {
        MoveHeroX(1);
        MoveHeroX(1);
    }
    if (event.key == "H" || event.key == "A" || (event.shiftKey && event.key == "ArrowLeft")) {
        MoveHeroX(-1);
        MoveHeroX(-1);
    }
    if (event.key == "K" || event.key == "W" || (event.shiftKey && event.key == "ArrowUp")) {
        MoveHeroY(-1);
        MoveHeroY(-1);
    }
    if (event.key == "J" || event.key == "S" || (event.shiftKey && event.key == "ArrowDown")) {
        MoveHeroY(1);
        MoveHeroY(1);
    }
    if (event.key == "l" || event.key == "d" || event.key == "ArrowRight") {
        MoveHeroX(1);
    }
    if (event.key == "h" || event.key == "a" || event.key == "ArrowLeft") {
        MoveHeroX(-1);
    }
    if (event.key == "k" || event.key == "w" || event.key == "ArrowUp") {
        MoveHeroY(-1);
    }
    if (event.key == "j" || event.key == "s" || event.key == "ArrowDown") {
        MoveHeroY(1);
    }

});

var timeOutFunctionId;

window.addEventListener('resize', function() {

    clearTimeout(timeOutFunctionId);

    timeOutFunctionId = setTimeout(window.location.reload(), 100000)
});

let canvs = {
    element: [],
    ctx: [],
    position: {
        x: [],
        y: []
    }
}


async function init() {
    let result = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
    const wasm = result.instance;
    go.run(wasm);

    start();
    heroStart();
}

function start() {
    CreateTheArena();
}
function heroStart() {
    CreateTheHero();
}

/*
 * @param {int} dimensionX
 * @param {int} dimensionY
*/
function Arena(dimesionX, dimesionY) {
    let container = document.getElementById("container");
    let dimension = 0;

    for (let i = 0; i < dimesionX; ++i) {

        for (let j = 0; j < dimesionY; ++j) {
            let theArena = document.createElement('canvas');
            theArena.innerHTML = "/";
            theArena.setAttribute('id', `arena${i}-${j}`);

            container.appendChild(theArena);

            canvs.position.x[dimension] = (42 * j);
            canvs.position.y[dimension] = (42 * i);

            ++dimension;
        }

        container.appendChild(document.createElement('br'));

    }

    let temp = container.querySelectorAll('canvas');

    for (let i = 0; i < dimesionX * dimesionY; i++) {
        canvs.element[i] = temp[i];
        canvs.ctx[i] = canvs.element[i].getContext("2d");
        // drawArena(canvs.ctx[i], canvs.position.x[i], canvs.position.y[i]);
        drawArena(canvs.ctx[i], 0, 0);
    }

}

function colorPath(ctx, color) {
    ctx.fillStyle = color;
    ctx.fillRect(0, 0, 420, 420);
    ctx.stroke();
}

function drawArena(ctx, x, y) {
    ctx.beginPath();
    ctx.fillStyle = "#181825";
    ctx.fillRect(x, y, 420, 420);
    ctx.stroke();
}

/*
 * @param {
 */
function Hero(theHero) {
    theHero.beginPath();
    theHero.arc(170, 75, 50, 0, 10);

    theHero.fillStyle = "#b4befe";
    theHero.fill();

    theHero.stroke();
}

function levelCompleate() {
    let text = "Level compleate!!\ndo you want to continue?"

    if (confirm(text)) {
        location.reload()
    } else {

    }
}

function colorPortal(ctx) {
    ctx.clearRect(60, 25, 130, 100);
}
