const go = new Go();
let wasm;
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {

    wasm = result.instance;
    go.run(wasm);
    console.log(wasm);

});

let canvs = {
    element: [],
    ctx: [],
    position: {
        x: [],
        y: []
    }
}

let theHero = {
    hero: null,
    ctxh: null,
    postion: {
        x: 0,
        y: 8
    }
};

function init() {
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
    console.log(x, y);
    ctx.beginPath();
    ctx.fillStyle = "#00FF00";
    ctx.fillRect(x, y, 420, 420);
    ctx.stroke();
}

function Hero() {
    theHero.hero = document.getElementById("hero");
    theHero.ctxh = theHero.hero.getContext("2d");

    theHero.ctxh.beginPath();
    theHero.ctxh.arc(170, 75, 73, 0, 10);
    theHero.ctxh.fillStyle = "#FF0000";
    theHero.ctxh.fill();
    theHero.ctxh.stroke();

    theHero.hero.style.left = theHero.postion.x + "px"
    theHero.hero.style.top = theHero.postion.y + "px"
}

function moveHero() {
    console.log("Entro");
    theHero.postion.x += 10
    theHero.hero.style.left = theHero.postion.x + "px";
}
