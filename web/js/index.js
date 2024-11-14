let canvs = {
    element: [],
    ctx: [],
    position: {
        x: [],
        y: []
    }
}

function init() {

    Arena();


    const hero = document.getElementById("hero");
    const ctxh = hero.getContext("2d");
    ctxh.beginPath();
    ctxh.fillStyle = "#FF0000";
    ctxh.fillRect(69, 69, 20, 20);
    ctxh.stroke();

}

function Arena() {
    let container = document.getElementById("container");
    let dimension = 0;

    for (let i = 0; i < 5; ++i) {
        for (let j = 0; j < 5; ++j) {
            container.innerHTML += '<canvas id="arena" tag="prov' + dimension + '>Sorry no cavs for you</canvas>';
            canvs.position.x[dimension] = (42 * j);
            canvs.position.y[dimension] = (42 * i);
            ++dimension;
        }
    }

    for (let i = 0; i < 25; i++) {
        canvs.element[i] = document.getElementById("arena");
        console.log(canvs.element[i]);
        canvs.ctx[i] = canvs.element[i].getContext("2d");
        drawArena(canvs.ctx[i], canvs.position.x[i], canvs.position.y[i]);
    }

}

function drawArena(ctx, x, y) {
    console.log(x, y, ctx);
    ctx.beginPath();
    ctx.fillStyle = "#00FF00";
    ctx.fillRect(x, y, 420, 420);
    ctx.stroke();
}

