document.addEventListener('DOMContentLoaded', async () => {
    const token = sessionStorage.getItem("token");
    let response;

    try {
        if (token) {
            response = await fetch("https://server-project-production-b671.up.railway.app/events/registered", {
                headers: {
                    "Authorization": `Bearer ${token}`
                }
            });
        } else {
            response = await fetch("https://server-project-production-b671.up.railway.app/events");
        }

        if (!response.ok) {
            console.error("Възникна грешка при рареждането на събитията")
        }

        const data = await response.json()
        console.dir(data)
    } catch (error) {
        console.error(error)
        alert("Възникна грешка при рареждането на събитията")
    }
});