const username = localStorage.getItem('username');

if (!username) {
    window.location.href = './index.html';
}

document.getElementById("nickname").innerHTML = username;

async function dataUpdate() {
    try {
        const response = await fetch("userdata.json", {
            method: 'GET',
            cache: "no-store"
        });

        if (!response.ok) {
            throw new Error("Ошибка сети: " + response.status);
        }

        const data = await response.json();
        const user = data[username];

        if (user) {
            console.log(user);
        } else {
            console.log('Пользователь не найден');
        }
    } catch (error) {
        console.error("Ошибка при загрузке данных", error);
    }
}

const updateInterval = setInterval(dataUpdate, 1000);

window.addEventListener('beforeunload', () => {
    clearInterval(updateInterval);
});

dataUpdate();
