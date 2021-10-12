const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const username = document.querySelector('#username')
const send = document.querySelector('#send')
const statusTyping = document.querySelector('#status')
const create = document.querySelector('#create')
const createInput = document.querySelector('.create-channel')

const protocol = window.location.protocol == 'https:' ? 'wss:' : 'ws:';
const host = window.location.host;
const path = window.location.pathname;

const url = `${protocol}//${host}${path}ws`;
const ws = new WebSocket(url);

const md = window.markdownit();

const wsSend = data => {
    ws.send(JSON.stringify({
        ...data,
        username: username.value,
    }));
};

const sendMessage = () => {
    wsSend({
        kind: 'message',
        content: input.value,
    });

    input.value = "";
};

/**
 * Update "<username> is typing..." messages
 */
// TODO: finish typing status
const updateTypingStatus = msg => {
    console.log("TEXTAREA CHANGED!")
    console.log('updateTyping got', msg);
    // Create a div object which will hold the message
    const message = document.createElement('div')

    // Set the attribute of the message div
    message.setAttribute('class', 'status-message')
    console.log("name: " + msg.username + " is typing")
    message.textContent = `${msg.username} is typing...`

    // Append the message to status div if size < 3
    if (statusTyping.childElementCount < 3) {
        statusTyping.appendChild(message)
    } else {
        console.log("Already 3 elements!")
    }
};

let typing = {
    isTyping: false,
    timer: null,
    timeout: 3000,

    /**
     * Sends a {kind: "typing"} message to the websocket
     * when the user has stopped typing.
     *
     * This function is debounced, so it only triggers after the user
     * has already stopped typing for a while, and not between every
     * keystroke
     */
    stopped() {
        this.isTyping = false;
        wsSend({
            kind: 'typing',
            isTyping: false,
        });
    },

    /**
     * Sends a {kind: "typing"} message to the websocket
     * when the user has started typing.
     *
     * Resets the timer for the stoppedTyping debounce every time it's
     * called, but only sends a websocket message if the user has
     * just started typing
     */
    started() {
        // If we don't get called again in 3 seconds (this.timeout ms),
        // we should call this.stopped();
        clearTimeout(this.timer);
        this.timer = setTimeout(() => { this.stopped(); }, this.timeout);

        if (!this.isTyping) {
            this.isTyping = true;
            wsSend({
                kind: 'typing',
                isTyping: true,
            });
        }
    },
};

// The user can send a message either by clicking "Send"
// or pressing "Enter" (unless shift is pressed).
send.addEventListener('click', sendMessage);
input.addEventListener('keyup', evt => {
    typing.started();
    if (evt.key === 'Enter' && !evt.shiftKey) {
        evt.stopPropagation();
        sendMessage();
    }
});

// User can create a new channel by typing a name and pressing the button
if (create) {
    create.addEventListener('click', () => {
        const channelName = createInput.value;
        window.open(`channel/${channelName}`, "_blank").focus();
    });
}

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

    // Scroll automatically
    messages.scrollTop = messages.scrollHeight
}

const handlers = {
    'message': insertMessage,
    'typing': updateTypingStatus,
};

ws.onmessage = ({ data }) => {
    const msg = JSON.parse(data);
    console.log('got', msg);

    // Returns if username or msg content is null
    if (msg.username == "" || msg.content == "") return;
    const handler = handlers[msg.kind];
    handler(msg);
};
