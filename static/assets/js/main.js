document.addEventListener('DOMContentLoaded', async () => {
    try {
        const response = await fetch("https://server-project-production-b671.up.railway.app/events");
        const data = await response.json();

        const list = document.getElementById('events-list');

        data.forEach(event => {
            const listItem = document.createElement('li');

            const eventTitle = document.createElement('h1');
            eventTitle.innerText = event.title;
            listItem.appendChild(eventTitle);

            const eventImage = document.createElement('img');
            eventImage.src = event.image_url;
            eventImage.alt = event.title;
            listItem.appendChild(eventImage);

            const eventDate = document.createElement('p');
            eventDate.innerText = `Дата: ${new Date(event.date).toLocaleDateString()}`;
            listItem.appendChild(eventDate);

            const eventAddress = document.createElement('p');
            eventAddress.innerText = `Място: ${event.address}`;
            listItem.appendChild(eventAddress);

            const eventDescription = document.createElement('p');
            eventDescription.innerText = event.description;
            listItem.appendChild(eventDescription);

            list.appendChild(listItem);
        });
    } catch (error) {
        console.error("Error fetching events:", error);
    }
});