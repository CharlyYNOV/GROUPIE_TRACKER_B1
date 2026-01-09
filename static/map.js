    //find "map" in the html
    const mapElement = document.getElementById('map');
    //retrieve information from the markers
    const rawData = mapElement.getAttribute('data-markers');
    
    // data recovery
    let markers = [];
    try {
        markers = JSON.parse(rawData);
    } catch (e) {
        console.error("Erreur récupération des données", e);
    }

   //map from leaflet
    var map = L.map('map').setView([20, 0], 2);
    
    // weird thing of leaflet 
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);
   
    //markers 
    if (Array.isArray(markers)) {
        markers.forEach(function(loc) {
            if (loc.lat !== 0) {
                L.marker([loc.lat, loc.lng])
                    .addTo(map)
                    .bindPopup(`<b>${loc.artist}</b><br>${cityDisplay}`);
            }
        });
    }
    