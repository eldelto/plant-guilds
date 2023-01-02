let selectedPlants = new Set();
let suggestions = new Array();
let goodCompanions = new Set();
let badCompanions = new Set();

function populateSuggestions(query) {
  query = query.toLowerCase();
  suggestions = new Array();
  plantGuilds.forEach((_, key) => {
    if (key.toLowerCase().indexOf(query) >= 0) {
      suggestions.push(key);
    }
  })
}

function updateSuggestions(query) {
  const list = window.document.getElementById("suggestions");
  list.innerHTML = "";
  suggestions.forEach((plantName) => {
    const li = document.createElement("li");
    li.onclick = () => selectPlant(plantName);
    li.textContent = plantName;
    list.appendChild(li);
  });


  if (query.length > 0) hideGoodCompanions()
  else showGoodCompanions();
}

function search() {
  const searchInput = window.document.getElementById("plant-search");
  populateSuggestions(searchInput.value);
  updateSuggestions(searchInput.value);
}

function updateSelectedPlants(query) {
  const list = window.document.getElementById("selected-plants"); // TODO: Move selection to on-page-load
  list.innerHTML = "";
  selectedPlants.forEach((plantName) => {
    const li = document.createElement("li");
    li.onclick = () => removePlant(plantName);
    li.textContent = plantName;
    list.appendChild(li);
  });

  updateCompanions();
}

function selectFirstPlant() {
  if (suggestions.length > 0) {
    const plantName = suggestions[0];
    selectPlant(plantName);
  }

  return false;
}

function selectPlant(name) {
  const searchInput = window.document.getElementById("plant-search");
  searchInput.value = "";
  search();
  selectedPlants.add(name);
  updateSelectedPlants();
}

function removePlant(name) {
  selectedPlants.delete(name);
  updateSelectedPlants();
}

function populateCompanions() {
  goodCompanions = new Set();
  badCompanions = new Set();

  selectedPlants.forEach((plantName) => {
    const plantInfo = plantGuilds.get(plantName);

    // Union
    goodCompanions = new Set([...goodCompanions, ...plantInfo.Good]);
    badCompanions = new Set([...badCompanions, ...plantInfo.Bad]);
  });

  for (var x of badCompanions) if (goodCompanions.has(x)) goodCompanions.delete(x);
  for (var x of selectedPlants) {
    if (goodCompanions.has(x)) goodCompanions.delete(x);
    if (badCompanions.has(x)) badCompanions.delete(x);
  }
}

function updateCompanions() {
  populateCompanions();

  const goodCompanionList = window.document.getElementById("good-companions"); // TODO: Move selection to on-page-load
  goodCompanionList.innerHTML = "";
  goodCompanions.forEach((plantName) => {
    const li = document.createElement("li");
    li.textContent = plantName;
    goodCompanionList.appendChild(li);
  });
}

function showGoodCompanions() {
  const goodCompanionList = window.document.getElementById("good-companions"); // TODO: Move selection to on-page-load
  goodCompanionList.removeAttribute("hidden");
}

function hideGoodCompanions() {
  const goodCompanionList = window.document.getElementById("good-companions"); // TODO: Move selection to on-page-load
  goodCompanionList.setAttribute("hidden", "true");
}