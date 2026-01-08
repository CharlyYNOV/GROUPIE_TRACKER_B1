
    const mapElement = document.getElementById('map');
    
    // recupérer données html
    const rawData = mapElement.getAttribute('data-markers');
    
    let markers = [];
    try {
        markers = JSON.parse(rawData);
    } catch (e) {
        console.error("Erreur de parsing des données concerts:", e);
    }

    // initialiser carte
    var map = L.map('map').setView([20, 0], 2);

    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);
    // marqueurs 
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
    