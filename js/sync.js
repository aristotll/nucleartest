async function aweek() {
    console.log('Monday')
    await playMusic() // block call
    console.log('Tuesday')
    await browseMoment()
    console.log('Wednesday')
}

async function playMusic() {
    await sleep(5000)
    console.log('paly music done')
}

async function browseMoment() {
    await sleep(5000)
    console.log('browse moment done')
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

aweek()
