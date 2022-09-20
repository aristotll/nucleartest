function aweek() {
    console.log('Monday')
    playMusic() // block call
    console.log('Tuesday')
    browseMoment()
    console.log('Wednesday')
}

function playMusic() {
    sleep(500)
}

function browseMoment() {
    sleep(500)
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}
