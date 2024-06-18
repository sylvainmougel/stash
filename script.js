const list = document.getElementById("stash");
let button = document.getElementById('button');

function showStash() {
    let totals = []
    let dates = []
    let textFieldValue = document.getElementById('myTextField').value;
    console.log(textFieldValue)
    const client = algoliasearch('1QMZVCS1V5', textFieldValue);
    const index = client.initIndex('stash');
    var bar = new Promise((resolve, reject) => {
        index.search('', {"hitsPerPage":300}).then(({hits}) => {
            hits.forEach((hit => {
                totals.push(hit["total"])
                dates.push(hit["date"])
            }))
            resolve()
        });
    })
    bar.then(() => {
            const trace1 = {
                x: dates,
                y: totals,
                text: dates,
                type: 'scatter'
            };
            const data = [trace1];
            Plotly.newPlot('plot', data);
        }
    )


}
showStash()

// Add a 'click' event listener to the button
button.addEventListener('click', function() {
    showStash()
});



