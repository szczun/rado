// Funkcja do odczytu danych z serwera
function fetchData() {
    fetch('/.netlify/functions/counter')
        .then(response => response.json())
        .then(data => {
            document.getElementById('count').textContent = data.count;
        })
        .catch(error => console.error('Błąd:', error));
}

// Funkcja do wysyłania zaktualizowanych danych na serwer
function incrementCounter() {
    fetch('/.netlify/functions/counter', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
    })
        .then(response => response.json())
        .then(data => {
            document.getElementById('count').textContent = data.count;
        })
        .catch(error => console.error('Błąd:', error));
}

// Inicjalizacja po załadowaniu strony
document.addEventListener('DOMContentLoaded', () => {
    fetchData();  // Pobierz początkowy stan

    // Obsługuje kliknięcie przycisku
    document.querySelector('.increment-button').addEventListener('click', () => {
        incrementCounter();  // Zwiększ licznik i zaktualizuj
    });
});

