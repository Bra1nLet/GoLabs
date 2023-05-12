
export async function getList() {
    let url = 'https://api.github.com/repos/javascript-tutorial/en.javascript.info/commits';
    let response = await fetch(url);

    let commits = await response.json(); // читаем ответ в формате JSON

    alert(commits[0].author.login);
}

export function setItem(item) {
    return fetch('127.0.0.1:8080/auth/test', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ item })
    })
        .then(data => data.json())
}