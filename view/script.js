const showoverlay = document.getElementById('login-overlay');
const showoverlay2 = document.getElementById('signup-overlay');

const login = () => {
    showoverlay.classList.toggle("hide");
}

const showsignup = () => {
    showoverlay.classList.toggle("hide");
    showoverlay2.classList.toggle("hide");
}

const hidesignup = () => {
    showoverlay2.classList.toggle("hide");
}


const handleLike = (table, id) => {
    let Body = {};
    Body.value = "Like";
    Body.PorM_id = id;
    Body.table = table;
    btn_like = document.getElementById('btn-like' + id);
    btn_dislike = document.getElementById('btn-dislike' + id);
    const like = document.getElementById("like" + id)
    const dislike = document.getElementById("dislike" + id)
    var request = fetch("/like", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(Body)
    })
        .then(response => {
            response.text().then((val) => {
                like.innerText = val[1]
                dislike.innerText = val[3]
                btn_like.classList.toggle("active");
                btn_dislike.classList.remove("active");
            })
        })
        .catch(error => alert("Erreur : " + error));
    console.log(request);

}

const handleDislike = (table, id) => {
    let Body = {};
    Body.value = "Dislike";
    Body.PorM_id = id;
    Body.table = table;
    console.log(Body);
    btn_like = document.getElementById('btn-like' + id);
    btn_dislike = document.getElementById('btn-dislike' + id);
    const like = document.getElementById("like" + id)
    const dislike = document.getElementById("dislike" + id)
    var request = fetch("/like", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(Body)
    })
        .then(response => {
            response.text().then((val) => {
                like.innerText = val[1]
                dislike.innerText = val[3]
                btn_like.classList.remove("active");
                btn_dislike.classList.toggle("active");
            })
        })
        .catch(error => alert("Erreur : " + error));
}

// Register Check

$("#confirm_password").keypress(function (e) {
    if (e.which == 13) {
        RegisterCheck();
    }
});

function RegisterCheck() {
    var Username = document.getElementById("inputUsername").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("Password-register").value;
    var confirm_password = document.getElementById("confirm_password").value;
    img = document.getElementById("imageUpload");
    var text = document.getElementById("register-text");
    if (password != confirm_password) {
        text.innerText = "Password doesn't match";
    } else {
        let Body = {};
        Body.Email = email;
        Body.Username = Username;
        Body.Password = password;
        if (img.files && img.files[0]) {
            var reader = new FileReader();
            reader.onload = function (e) {
                Body.Image = e.target.result;
                var text_login = document.getElementById("login_text");
                text_login.innerText = Body.Image;
                var body = document.getElementById("Body");
                var btn = document.getElementById("btn-login");
                body.style.cursor = "wait";
                btn.style.cursor = "wait";
                var request = fetch("/register", {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(Body)
                })
                    .then(response => {
                        response.text().then((val) => {
                            if (val == "-1") {
                                text.innerText = "Email or username already used";
                                body.style.cursor = "default";
                                btn.style.cursor = "default";
                            } else if (val == "-2") {
                                text.innerText = "Not enough information";
                                body.style.cursor = "default";
                                btn.style.cursor = "default";
                            } else if (val == "1") {
                                var text_login = document.getElementById("login_text");
                                body.style.cursor = "default";
                                btn.style.cursor = "default";
                                showsignup();
                                text_login.innerText = "You are now registered";
                            }
                        })
                    })
                    .catch(error => alert("Erreur : " + error));
            }
            reader.readAsDataURL(img.files[0]);
        } else {
            var body = document.getElementById("Body");
            var btn = document.getElementById("btn-login");
            body.style.cursor = "wait";
            btn.style.cursor = "wait";
            var request = fetch("/register", {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(Body)
            })
                .then(response => {
                    response.text().then((val) => {
                        if (val == "-1") {
                            text.innerText = "Email or username already used";
                            body.style.cursor = "default";
                            btn.style.cursor = "default";
                        } else if (val == "-2") {
                            text.innerText = "Not enough information";
                            body.style.cursor = "default";
                            btn.style.cursor = "default";
                        } else if (val == "1") {
                            var text_login = document.getElementById("login_text");
                            body.style.cursor = "default";
                            btn.style.cursor = "default";
                            showsignup();
                            text_login.innerText = "You are now registered";
                        }
                    })
                })
                .catch(error => alert("Erreur : " + error));
            reader.readAsDataURL(img.files[0]);
        }
    }
}



// Login Check

$("#Password-login").keypress(function (e) {
    if (e.which == 13) {
        Login_check();
    }
});

function Login_check() {
    var Username = document.getElementById("username").value;
    var password = document.getElementById("Password-login").value;
    var text = document.getElementById("login_text");
    var body = document.getElementById("Body");
    var btn = document.getElementById("btn-login");
    body.style.cursor = "wait";
    btn.style.cursor = "wait";
    var request = fetch("/login", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            Username: Username,
            password: password
        })
    })
        .then(response => {
            response.text().then((val) => {
                if (val == "-1") {
                    text.innerText = "Wrong Username or Password";
                    body.style.cursor = "default";
                    btn.style.cursor = "default";
                } else if (val == "1") {
                    location.reload();
                }
            })
        })
        .catch(error => alert("Erreur : " + error));
}
