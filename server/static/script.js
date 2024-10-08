const login_button = document.querySelector(".button");
const reg_button = document.querySelector(".reg_button");
const closereg_button = document.querySelector(".closereg");
const reg_menu = document.getElementById('reg_menu');
const confirm_button = document.getElementById("confirm");

reg_menu.style.display = 'none';
reg_menu.style.opacity = '0';
reg_menu.style.transition = 'opacity 0.3s';

login_button.addEventListener('click', async function () {
    var username = document.querySelector("#login_id").value;
    localStorage.setItem('username', username);
    var password = document.querySelector("#pass_id").value;
    const wrongpw = document.getElementById('wrongPW');

    try {
        const response = await fetch("/userdata.json");
        if (!response.ok) {
            throw new Error('Ошибка сети: ' + response.status);
        }

        const data = await response.json();
        const user = data[username];

        if (user) {
            if (user.password === password) {
                history.pushState(null, '', './login.html');
                window.location.href = './login.html';
            } else {
                wrongpw.style.display = 'flex';
                setTimeout(() => {
                    wrongpw.style.display = 'none';
                }, 2000);
            }
        } else {
            console.log('Пользователь не найден');
        }

    } catch (error) {
        console.error('Ошибка при загрузке данных', error);
    }
});

reg_button.addEventListener('click', function () {
    reg_menu.style.display = 'flex';
    setTimeout(() => {
        reg_menu.style.opacity = '1';
    }, 10);
});

closereg_button.addEventListener('click', function () {
    reg_menu.style.opacity = '0';
    setTimeout(() => {
        reg_menu.style.display = 'none';
    }, 300);
});

confirm_button.addEventListener('click', async function () {
    var username = document.querySelector("#reg-login_id").value;
    var pass1 = document.querySelector("#reg-pass_id").value;
    var pass2 = document.querySelector("#reg-pass2_id").value;

    const new_user = {
        login: username,
        password: pass1,
        signals: []
    };

    if (username == "")
        console.log('1');
    else if (pass1 != pass2)
        console.log('2');
    else {
        try {
            const response = await fetch("./userdata.json", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(new_user)
            });

            if (response.ok) {
                console.log('Пользователь успешно добавлен');
            } else {
                console.log('Ошибка при добавлении пользователя');
            }

        } catch (error) {
            console.log("Ошибка ", error);
        }
    }
});
