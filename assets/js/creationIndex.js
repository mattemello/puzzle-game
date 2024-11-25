function init() {
    let col = Math.floor(window.innerWidth / 50);
    let raw = Math.floor(window.innerHeight / 50);

    let contain = document.getElementById("bg-canvas")

    for (let i = 0; i < col; i++) {
        for (let j = 0; j < raw; j++) {

            let can = '<canvas style="border: 1px solid black;">nocanvasforyou</canvas>';

            contain.innerHTML += can;

        }
    }

    let long = contain.querySelectorAll("canvas")

    for (let i = 0; i < (col - 2) * raw; i++) {

        let ctx = long[i].getContext("2d")

        colorMouseleave(ctx)

        long[i].addEventListener("mouseenter", function() {
            colorMouseover(ctx)
        });
        long[i].addEventListener("mouseleave", function() {
            colorMouseleave(ctx);
        });

    }

}

function colorMouseleave(ctx) {
    ctx.beginPath();
    ctx.fillStyle = "#181825";
    ctx.fillRect(0, 0, 420, 420);
    ctx.stroke();

}

function colorMouseover(ctx) {
    ctx.beginPath();
    ctx.fillStyle = "#182936";
    ctx.fillRect(0, 0, 420, 420);
    ctx.stroke();

}
