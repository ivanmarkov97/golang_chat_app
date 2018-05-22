var websocket = new WebSocket("ws://127.0.0.1:8080/ws");

var message_container = document.getElementById("message_container");
var my_text = document.getElementById("text");
var my_name = document.getElementById("my_name");

websocket.onopen = function() {
	;
};

websocket.onclose = function() {
	;
};


websocket.onmessage = function(event) {
	//alert("Data :" + event.data);
	incoming_message = JSON.parse(event.data);

	var new_message_block = document.createElement('div');
	new_message_block.className = 'message_block';

	var new_message_text = document.createElement('div');
	new_message_text.className = 'message_text';
	new_message_text.innerHTML = incoming_message["message"];
	//new_message_text.innerHTML = incoming_message;

	var new_message_author = document.createElement('div');
	new_message_author.className = 'author';
	new_message_author.innerHTML = "@" + incoming_message["username"];

	new_message_block.appendChild(new_message_text);
	new_message_block.appendChild(new_message_author);
	message_container.appendChild(new_message_block);

	message_container.scrollTop = message_container.scrollHeight;

	console.log(event.data);
	console.log(incoming_message);
};

websocket.onerror = function(error) {
	alert("Error " + error.message);
};

document.getElementById("send_btn").addEventListener("click", function(){
    var my_message = my_text.value;
    websocket.send(JSON.stringify({"email": "hey@mail.com", "username": my_name.value, "message": my_message}));
    my_text.value = "";
});
