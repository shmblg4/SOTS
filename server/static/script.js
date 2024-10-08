const login_button = document.querySelector(".button");
const reg_button = document.querySelector(".reg_button");
const closereg_button = document.querySelector(".closereg");
const reg_menu = document.getElementById('reg_menu');
const confirm_button = document.getElementById("confirm");
const wrongpw = document.getElementById('wrongPW');

reg_menu.style.display = 'none';
reg_menu.style.opacity = '0';
reg_menu.style.transition = 'opacity 0.3s';

login_button.addEventListener('click', async function () {
    const username = document.querySelector("#login_id").value;
    const password = document.querySelector("#pass_id").value;
    localStorage.setItem('username', username);

    try {
        const response = await fetch("/request", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                action: "login",
                user: {
                    login: username,
                    password: password
                }
            })
        });

        if (response.ok) {
            // Успешный вход
            history.pushState(null, '', './login.html');
            window.location.href = './login.html';
        } else {
            // Ошибка аутентификации
            wrongpw.style.display = 'flex';
            setTimeout(() => {
                wrongpw.style.display = 'none';
            }, 2000);
        }

    } catch (error) {
        console.error('Ошибка при аутентификации', error);
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

    if (username === "")
        console.log('Введите имя пользователя');
    else if (pass1 !== pass2)
        console.log('Пароли не совпадают');
    else {
        const new_user = {
            login: username,
            password: pass1
        };

        try {
            const response = await fetch("/request", { // Измените на корректный маршрут
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    action: "register",
                    user: new_user
                })
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
