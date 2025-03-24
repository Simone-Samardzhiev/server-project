document.getElementById("login-button").addEventListener("click", async (event) => {
    event.preventDefault()

    const email = document.getElementById("email-login").value
    const password = document.getElementById("password-login").value
    const errorText = document.getElementById("login-error-message")

    try {
        const response = await fetch("https://server-project-production-b671.up.railway.app/users/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({email, password})
        })

        switch (response.status) {
            case 201:
                const body = await response.json()
                const token = body.token;
                document.cookie = `token=${token}; path=/; Secure; HttpOnly; SameSite=Strict`;
                errorText.style.display = "none"
                break;
            default:
                errorText.textContent = "Възникна грешка"
                errorText.style.display = "block"
                console.error(await response.json())
        }
    } catch (error) {
        console.log(error)
    }
})