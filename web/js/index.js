let position = {
    x: 0,
    y: 0
};

const dimension = {
    rows: 10,
    collo: 10
}

let ctx = [];

function init() {
    let cont = document.getElementById("tableGame");
    let position = 0

    for (let i = 0; i < dimension.rows; i++) {
        cont.innerHTML += "<tr>"

        for (let j = 0; j < dimension.collo; j++) {
            cont.innerHTML += '<th><canvas class="canvas" id="canvas" onmouseleave="backColor(' + position + ')" onmouseenter="changeColor(' + position + ')">Sorry, bet no canvas</canvas></th>';
            ++position;
        }

        cont.innerHTML += "</tr>"
    }

    const canvas = document.querySelectorAll("canvas");

    for (let i = 0; i < dimension.rows * dimension.collo; i++) {
        ctx[i] = canvas[i].getContext("2d");
        if (i === 0) {
            drawnChar(i, 0, 0);
        }
        square(ctx[i]);
    }
}

function drawnChar(i, x, y) {
    console.log("enter");
    ctx[i].rect(x, y, 100, 100);
    ctx[i].lineWidth = 1;
    ctx[i].stroke();
    ctx[i].fillStyle = "black";
    ctx[i].fill();
}

function square(ctx) {
    ctx.rect(position.x, position.y, 300, 300);
    ctx.lineWidth = 1;
    ctx.stroke();
    ctx.fillStyle = "teal";
    ctx.fill();
}

function backColor(position) {
    ctx[position].fillStyle = "teal";
    ctx[position].fill();
}

function changeColor(position) {
    ctx[position].fillStyle = "black";
    ctx[position].fill();
}

function mov() {
    position.x += 10;
    position.y += 1;

    for (i = 0; i < 6; i++) {
        square(ctx[i]);
    }
}

