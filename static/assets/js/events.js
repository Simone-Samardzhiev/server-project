const getCookie = (name) => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

document.addEventListener('DOMContentLoaded', async () => {
    try {
        const token = getCookie('token');
        const endpoint = token ? 'https://server-project-production-b671.up.railway.app/events/registered'
            : 'https://server-project-production-b671.up.railway.app/events';

        const headers = token ? { 'Authorization': `Bearer ${token}` } : {};

        const response = await fetch(endpoint, { headers });

        if (!response.ok) throw new Error('Грешка при зареждане на събития');
        const events = await response.json();

        const list = document.getElementById('events-list');
        list.innerHTML = '';

        events.forEach(event => {
            const listItem = document.createElement('li');

            if (token) {
                const checkboxContainer = document.createElement('div');
                checkboxContainer.className = 'checkbox-container';

                const checkbox = document.createElement('input');
                checkbox.type = 'checkbox';
                checkbox.checked = event.is_registered || false;
                checkbox.disabled = true; // Remove when registration API is ready
                checkbox.className = 'registration-checkbox';

                // TODO: Add registration logic when backend is ready
                checkbox.addEventListener('change', () => {
                    console.log(`Registration changed for event ${event.id} to ${checkbox.checked}`);
                    // Implement updateRegistration() when API is ready
                });

                const label = document.createElement('label');
                label.textContent = ' Регистриран съм';
                label.prepend(checkbox);

                checkboxContainer.appendChild(label);
                listItem.appendChild(checkboxContainer);
            }

            const elements = {
                title: ['h1', event.title],
                image: ['img', null, { src: event.image_url, alt: event.title }],
                date: ['p', `Дата: ${new Date(event.date).toLocaleDateString()}`],
                address: ['p', `Място: ${event.address}`],
                description: ['p', event.description]
            };

            Object.entries(elements).forEach(([key, [tag, text, attributes]]) => {
                const el = document.createElement(tag);
                if (text) el.textContent = text;
                if (attributes) Object.entries(attributes).forEach(([name, value]) => el[name] = value);
                listItem.appendChild(el);
            });

            list.appendChild(listItem);
        });

    } catch (error) {
        console.error("Error:", error);
        alert('Грешка при зареждане на събитията');
    }
});