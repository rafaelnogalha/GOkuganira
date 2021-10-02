const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const username = document.querySelector('#username')
const send = document.querySelector('#send')
const statusTyping = document.querySelector('#status')

const protocol = window.location.protocol == 'https:' ? 'wss:' : 'ws:';
const url = `${protocol}//${window.location.host}/ws`;
const ws = new WebSocket(url);

const sendMessage = () => {
	const message = {
		kind: 'message',
		username: username.value,
		content: input.value,
	}

	ws.send(JSON.stringify(message));
	input.value = "";
};

/**
 * Update "<username> is typing..." messages
 */

// TODO: finish typing status
const updateTyping = msg => {
	console.log("TEXTAREA CHANGED!")
	console.log('updateTyping got', msg);
	// Create a div object which will hold the message
	const message = document.createElement('div')

	// Set the attribute of the message div
	message.setAttribute('class', 'status-message')
	console.log("name: " + msg + " is typing")
	message.textContent = `${msg} is typing...`

	// Append the message to our chat div
	statusTyping.appendChild(message)
};

// The user can send a message either by clicking "Send"
// or pressing "Enter" (unless shift is pressed).
send.addEventListener('click', sendMessage);
input.addEventListener('input', updateTyping);
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
const insertMessage = messageObj => {
	// Create a div object which will hold the message
	const message = document.createElement('div')

	// Set the attribute of the message div
	console.log("name: " + messageObj.username + " content: " + messageObj.content)
	message.textContent = `${messageObj.username}: ${messageObj.content}`
	message.setAttribute('class', 'chat-message')

	// Append the message to our chat div
	messages.appendChild(message)

	// Scroll automatically
	messages.scrollTop = messages.scrollHeight
}

const handlers = {
	'message': insertMessage,
	'typing': updateTyping,
};

ws.onmessage = ({ data }) => {
	const msg = JSON.parse(data);
	console.log('got', msg);
	
	// Returns if username or msg content is null
	if (msg.username == "" || msg.content == "") return;
	const handler = handlers[msg.kind];
	handler(msg);
};
