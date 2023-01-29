const downButton = document.getElementById('volDownButton');
const upButton = document.getElementById('volUpButton');
const prevButton = document.getElementById('previousButton');
const playButton = document.getElementById('playPauseButton');
const nextButton = document.getElementById('nextButton');

// Percentage
const step = 10

function adjustVol(amount, direction) {
    console.log('adjust volume: ', amount);
    fetch('/vol', {method: 'POST', body: JSON.stringify({direction: direction, amount: amount})})
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
    console.log('previous button was clicked');
    fetch('/previous', {method: 'POST'})
        .then(function(res) {
            if(res.ok) {
                console.log('click was Good');
                return;
            }
            throw new Error('Request failed.')
        })
        .catch(function(error) {
            console.log(error)
        })
});

playButton.addEventListener('click', function(e) {
    console.log('play button was clicked');
    fetch('/playPause', {method: 'POST'})
        .then(function(res) {
            if(res.ok) {
                console.log('click was Good');
                return;
            }
            throw new Error('Request failed.')
        })
        .catch(function(error) {
            console.log(error)
        })
});

nextButton.addEventListener('click', function(e) {
    console.log('next button was clicked');
    fetch('/next', {method: 'POST'})
        .then(function(res) {
            if(res.ok) {
                console.log('click was Good');
                return;
            }
            throw new Error('Request failed.')
        })
        .catch(function(error) {
            console.log(error)
        })
});
