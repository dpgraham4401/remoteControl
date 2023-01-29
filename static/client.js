const downButton = document.getElementById('volDownButton');
const upButton = document.getElementById('volUpButton');
const prevButton = document.getElementById('previousButton');
const playButton = document.getElementById('playPauseButton');
const nextButton = document.getElementById('nextButton');

// Percentage
const step = 10

function adjustVol(amount, direction) {
    fetch('/vol', {method: 'POST', body: JSON.stringify({direction: direction, amount: amount})})
        .then(function(res) {
        })
        .catch(function(error) {
            console.warn('error', error)
        })
}

function playCommand(playerArg) {
    fetch('/play', {method: 'POST', body: JSON.stringify({command: playerArg})})
        .then(function(res) {
        })
        .catch(function(error) {
            console.warn('error', error)
        })
}

downButton.addEventListener('click', function(e) {
    adjustVol(5, "down")
});

upButton.addEventListener('click', function(e) {
    adjustVol(5, "up")
});

prevButton.addEventListener('click', function(e) {
    playCommand("previous")
});

playButton.addEventListener('click', function(e) {
    playCommand("play-pause")
});

nextButton.addEventListener('click', function(e) {
    playCommand("next")
});
