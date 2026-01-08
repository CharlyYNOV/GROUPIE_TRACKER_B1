document.addEventListener('DOMContentLoaded', function() {
    const mapElement = document.getElementById('map');
    
    // On récupère la chaîne JSON stockée dans le HTML
    const rawData = mapElement.getAttribute('data-markers');
    
    let markers = [];
    try {
        markers = JSON.parse(rawData);
    } catch (e) {
        console.error("Erreur de parsing des données concerts:", e);
    }

    // Initialisation de la carte
    var map = L.map('map').setView([20, 0], 2);

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; OpenStreetMap contributors'
    }).addTo(map);

    // Ajout des marqueurs
    if (Array.isArray(markers)) {
        markers.forEach(function(loc) {
            if (loc.lat !== 0) {
                var cityDisplay = loc.city.replace(/_/g, ' ').replace(/-/g, ', ');
                L.marker([loc.lat, loc.lng])
                    .addTo(map)
                    .bindPopup(`<b>${loc.artist}</b><br>${cityDisplay}`);
            }
        });
    }
});