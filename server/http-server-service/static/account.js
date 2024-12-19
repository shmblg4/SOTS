// account.js

const username = localStorage.getItem('username');


if (!username) {
    window.location.href = './index.html';
}

document.getElementById("nickname").innerHTML = username;
const signals = document.getElementById("user-signals");
async function dataUpdate() {
    try {
        const response = await fetch(`/loadData?username=${encodeURIComponent(username)}`, {
            method: 'GET',
            cache: "no-store"
        });

        if (!response.ok) {
            throw new Error("Ошибка сети: " + response.status);
        }

        const data = await response.json();
        if (data) {
            signals.innerHTML = data[data.length - 1];
        } else {
            console.log('Данные не найдены');
        }
    } catch (error) {
        console.error("Ошибка при загрузке данных, %s", error);
    }
}

const updateInterval = setInterval(dataUpdate, 1000);

window.addEventListener('beforeunload', () => {
    clearInterval(updateInterval);
});

dataUpdate();