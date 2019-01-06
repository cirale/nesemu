const canvas = document.getElementById("emu");
var ctx = canvas.getContext('2d');
const ws = new WebScoket("ws://localhost:18080/");

ws.onmessage = event => {
    var imageData = ctx.getImageData(0,0,canvas.width, canvas.height);
    var data = imageData.data
    
}
