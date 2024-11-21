const go = new Go();
//wasm.CreateTheArena();

//TODO: use the map for mapping the key (←↓↑→) or (wasd)


document.addEventListener('keydown', function (event) {
    if (event.key == "l") {
        MoveHeroX(1);
    }
    if (event.key == "h") {
        MoveHeroX(-1);
    }
    if (event.key == "k") {
        MoveHeroY(-1);
    }
    if (event.key == "j") {
        MoveHeroY(1);
    }
    if (event.key == "L") {
        MoveHeroX(1);
        MoveHeroX(1);
    }
    if (event.key == "H") {
        MoveHeroX(-1);
        MoveHeroX(-1);
    }
    if (event.key == "K") {
        MoveHeroY(-1);
        MoveHeroY(-1);
    }
    if (event.key == "J") {
        MoveHeroY(1);
        MoveHeroY(1);
    }

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

function Arena(dimesionX, dimesionY) {
    let container = document.getElementById("container");
    let dimension = 0;

    for (let i = 0; i < dimesionX; ++i) {

        for (let j = 0; j < dimesionY; ++j) {
            let theArena = document.createElement('canvas');
            theArena.innerHTML = "Sorry no cavs for you";
            theArena.setAttribute('id', `arena${i}${j}`);

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

function Hero(theHero) {
    theHero.beginPath();
    theHero.arc(170, 75, 50, 0, 10);

    theHero.fillStyle = "#b4befe";
    theHero.fill();

    theHero.stroke();
}

