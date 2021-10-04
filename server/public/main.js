const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const username = document.querySelector('#username')
const send = document.querySelector('#send')

const protocol = window.location.protocol == 'https:' ? 'wss:' : 'ws:';
const url = `${protocol}//${window.location.host}/ws`;
const ws = new WebSocket(url);

const md = window.markdownit();

const sendMessage = () => {
    const message = {
        kind: 'message',
        username: username.value,
        content: input.value,
    }

    ws.send(JSON.stringify(message));
    input.value = "";
};

// The user can send a message either by clicking "Send"
// or pressing "Enter" (unless shift is pressed).
send.addEventListener('click', sendMessage);
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
    message.setAttribute('class', 'chat-message')
    console.log("name: " + messageObj.username + " content: " + messageObj.content)
    message.innerHTML = md.renderInline(`${messageObj.username}: ${messageObj.content}`)

    // Append the message to our chat div
    messages.appendChild(message)

    // Insert the message as the first message of our chat
    messages.insertBefore(message, messages.firstChild)
}

/**
 * Update "<username> is typing..." messages
 */
const updateTyping = msg => {
    console.log('updateTyping got', msg);
};

const handlers = {
    'message': insertMessage,
    'typing': updateTyping,
};

ws.onmessage = ({ data }) => {
    const msg = JSON.parse(data);
    console.log('got', msg);
    const handler = handlers[msg.kind];
    handler(msg);
};
