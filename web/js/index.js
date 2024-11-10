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
    let cont = document.getElementById("container");
    let position = 0

    for (let i = 0; i < dimension.rows; i++) {
        for (let j = 0; j < dimension.collo; j++) {
            cont.innerHTML += '<canvas class="canvas" id="canvas" onmouseleave="backColor(' + position + ')" onmouseenter="changeColor(' + position + ')">Sorry, bet no canvas</canvas>';
            ++position;
        }
    }

    const canvas = document.querySelectorAll("canvas");

    for (let i = 0; i < dimension.rows * dimension.collo; i++) {
        ctx[i] = canvas[i].getContext("2d");
        square(ctx[i]);
    }
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

