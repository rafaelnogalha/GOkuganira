const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const username = document.querySelector('#username')
const send = document.querySelector('#send')

const protocol = window.location.protocol == 'https:' ? 'wss:' : 'ws:';
const url = `${protocol}//${window.location.host}/ws`;
const ws = new WebSocket(url);

wss.onmessage = function (msg) {
	console.log(msg.data)
    insertMessage(JSON.parse(msg.data))
};

const sendMessage = () => {
    const message = {
		username: username.value,
		content: input.value,
    }

    wss.send(JSON.stringify(message));
    input.value = "";
};

send.onclick = sendMessage;
input.addEventListener('keyup', evt => {
	if (evt.key === 'Enter' && !evt.shiftKey) {
		evt.stopPropagation();
		sendMessage();
	}
});


/**
 * Insert a message into the UI
 * @param {Message that will be displayed in the UI} messageObj
 */
function insertMessage(messageObj) {
	// Create a div object which will hold the message
	const message = document.createElement('div')

	// Set the attribute of the message div
	message.setAttribute('class', 'chat-message')
	console.log("name: " + messageObj.username + " content: " + messageObj.content)
	message.textContent = `${messageObj.username}: ${messageObj.content}`

	// Append the message to our chat div
	messages.appendChild(message)

	// Insert the message as the first message of our chat
	messages.insertBefore(message, messages.firstChild)
}
