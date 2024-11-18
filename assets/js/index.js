const go = new Go();
//wasm.CreateTheArena();


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
            theArena.setAttribute('id', 'arena');

            container.appendChild(theArena);

            canvs.position.x[dimension] = (42 * j);
            canvs.position.y[dimension] = (42 * i);

            ++dimension;
        }

        container.appendChild(document.createElement('br'));

    }

    let temp = container.querySelectorAll('canvas');

    for (let i = 0; i < 25; i++) {
        canvs.element[i] = temp[i];
        canvs.ctx[i] = canvs.element[i].getContext("2d");
        // drawArena(canvs.ctx[i], canvs.position.x[i], canvs.position.y[i]);
        drawArena(canvs.ctx[i], 0, 0);
    }

}

function drawArena(ctx, x, y) {
    ctx.beginPath();
    ctx.fillStyle = "#00FF00";
    ctx.fillRect(x, y, 420, 420);
    ctx.stroke();
}

function Hero(theHero) {

    theHero.beginPath();
    theHero.arc(170, 75, 73, 0, 10);
    theHero.fillStyle = "#FF0000";
    theHero.fill();
    theHero.stroke();

}

/*function moveHero() {
    console.log("Entro");
    theHero.postion.x += 10
    theHero.hero.style.left = theHero.postion.x + "px";
}*/
