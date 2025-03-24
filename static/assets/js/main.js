document.addEventListener('DOMContentLoaded', () => {
    fetch("https://server-project-production-b671.up.railway.app/events")
        .then(response => response.json())
        .then(data => {
            let list = document.getElementById('events-list'); // Corrected ID to 'events-list'
            data.forEach(event => {
                let listItem = document.createElement('li'); // Create the <li> element

                let eventTitle = document.createElement('h1');
                eventTitle.innerText = event.title;

                listItem.appendChild(eventTitle);

                let eventImage = document.createElement('img');
                eventImage.src = event.image_url;
                eventImage.alt = event.title;

                listItem.appendChild(eventImage);

                let eventDate = document.createElement('p');
                eventDate.innerText = `Дата: ${new Date(event.date).toLocaleDateString()}`;

                let eventAddress = document.createElement('p');
                eventAddress.innerText = `Място: ${event.address}`;

                let eventDescription = document.createElement('p');
                eventDescription.innerText = event.description;

                listItem.appendChild(eventDate);
                listItem.appendChild(eventAddress);
                listItem.appendChild(eventDescription);

                list.appendChild(listItem);
            });
        })
        .catch(error => {
            console.log(error);
        });
});