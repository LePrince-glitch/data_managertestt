document.addEventListener("DOMContentLoaded", function() {
  const sidebarLinks = document.querySelectorAll(".sidebar a");

  sidebarLinks.forEach(function(link) {
    link.addEventListener("click", function(event) {
      event.preventDefault();

      const targetId = this.getAttribute("href").substring(1);
      const targetSection = document.getElementById(targetId);

      if (targetSection) {
        // Masquer toutes les sections
        document
          .querySelectorAll(".content section")
          .forEach(function(section) {
            section.classList.remove("active");
          });

        // Afficher uniquement la section ciblée
        targetSection.classList.add("active");
      }
    });
  });
});

document
  .getElementById("modifyButton")
  .addEventListener("click", function(event) {
    event.preventDefault();
    toggleFields();
  });

function toggleFields() {
  var fields = document.querySelectorAll("#myForm input");
  var modifyButton = document.getElementById("modifyButton");

  if (modifyButton.textContent === "Modifier") {
    fields.forEach(function(field) {
      field.classList.remove("disabled");
    });
    modifyButton.textContent = "Annuler";
  } else {
    fields.forEach(function(field) {
      field.classList.add("disabled");
    });
    modifyButton.textContent = "Modifier";
  }
}

const table = document.getElementById("myTable");
const formContainer = document.getElementById("form-container");
const form = document.getElementById("myForm");
const siteInput = document.getElementById("site");
const numeroInput = document.getElementById("numero");
const clientInput = document.getElementById("client");
const volumesInput = document.getElementById("volumes");
const rechargeInput = document.getElementById("recharge");
const expirationInput = document.getElementById("expiration");
const Btnfermer = document.getElementById("Btnfermer");
const customerId = document.getElementById("customer-id");
const siteId = document.getElementById("site-id");
const numberId = document.getElementById("number-id");
const rechargeId = document.getElementById("recharge-id");

table.addEventListener("click", function(event) {
  const clickedRow = event.target.parentNode;
  const cells = clickedRow.querySelectorAll("td");

  siteInput.value = cells[1].textContent;
  numeroInput.value = cells[2].textContent;
  clientInput.value = cells[3].textContent;
  volumesInput.value = cells[4].textContent;
  rechargeInput.value = cells[5].textContent;
  expirationInput.value = cells[6].textContent;
  formContainer.classList.add("show");
});

Btnfermer.addEventListener("click", function() {
  formContainer.classList.remove("show");
});

document.addEventListener("DOMContentLoaded", function() {
  var filtresBtn = document.getElementById("filtresBtn");
  var autres = document.getElementById("autres");

  filtresBtn.addEventListener("click", function() {
    autres.style.display = autres.style.display === "flex" ? "none" : "flex";
  });
});

// **********************USER FORMULAIRE

document.addEventListener("DOMContentLoaded", function() {
  var cliqueUser = document.getElementById("cliqueUser");
  var formulaireU = document.getElementById("formulaireU");
  var closeBtnU = document.getElementById("closeBtnU");
  cliqueUser.addEventListener("click", function() {
    formulaireU.style.display =
      formulaireU.style.display === "flex" ? "none" : "flex";
  });
  closeBtnU.addEventListener("click", function() {
    formulaireU.style.display = "none";
  });
});

/***************** espace dossier******************************* */
document.addEventListener("DOMContentLoaded", function() {
  var filtresBtnd = document.getElementById("filtresBtnd");
  var autres1 = document.getElementById("autres1");

  filtresBtnd.addEventListener("click", function() {
    autres1.style.display = autres1.style.display === "flex" ? "none" : "flex";
  });
});

//**********gestion du formulaire de modification des dossiers

// Fonction pour afficher le formulaire de modification
function afficherFormulaire() {
  document.getElementById("modificationForm").classList.remove("hidden");
  document.getElementById("modificationForm").classList.add("visible");
}
// Fonction pour cacher le formulaire de modification
function cacherFormulaire() {
  document.getElementById("modificationForm").classList.remove("visible");
  document.getElementById("modificationForm").classList.add("hidden");
}
// Fonction pour remplir le formulaire avec les valeurs par défaut
function remplirFormulaire(
  nomSociete,
  reference,
  nomCommercial,
  secteur,
  contact,
  statut
) {
  document.getElementById("inputNomSociete").value = nomSociete;
  document.getElementById("inputReference").value = reference;
  document.getElementById("inputNomCommercial").value = nomCommercial;
  document.getElementById("inputSecteur").value = secteur;
  document.getElementById("inputContact").value = contact;
  document.getElementById("inputStatut").value = statut;
}

// Événement de clic sur une boîte
document.querySelectorAll(".boxx").forEach((box) => {
  box.addEventListener("click", () => {
    // Remplacer les valeurs par défaut du formulaire avec les valeurs des span correspondants
    remplirFormulaire(
      box.querySelector("#nomsociete").innerText,
      box.querySelector("#Reference").innerText,
      box.querySelector("#commercial").innerText,
      box.querySelector("#Secteur").innerText,
      box.querySelector("#contact").innerText,
      box.querySelector("#statut").innerText
    );
    // Afficher le formulaire de modification
    afficherFormulaire();
  });
});
document.getElementById("closeBtnMD").addEventListener("click", () => {
  cacherFormulaire();
});
// Événement de clic sur le bouton Enregistrer
document.getElementById("enregistrer").addEventListener("click", () => {
  // Code pour enregistrer les modifications
});

// Événement de clic sur le bouton Modifier
document.getElementById("modifier").addEventListener("click", () => {
  // Code pour modifier le dossier
});

// Événement de clic sur le bouton Supprimer
document.getElementById("supprimer").addEventListener("click", () => {
  // Code pour supprimer le dossier
});

//**********tableau dossier

// Sélectionnez toutes les lignes de tableau
var rows = document.querySelectorAll("tr");

// Parcourez chaque ligne de tableau
rows.forEach(function(row) {
  // Ajoutez un écouteur d'événements 'click' à chaque ligne de tableau
  row.addEventListener("click", function() {
    // Supprimez la classe 'selected' de toutes les lignes de tableau
    rows.forEach(function(row) {
      row.classList.remove("selected");
    });
    // Ajoutez la classe 'selected' à la ligne de tableau cliquée
    row.classList.add("selected");
  });
});
//****** formulaire dossier

document.addEventListener("DOMContentLoaded", function() {
  var newdossier = document.getElementById("newdossier");
  var ajoutclient = document.getElementById("ajoutclient");
  var closeBtn = document.getElementById("closeBtn");

  newdossier.addEventListener("click", function() {
    ajoutclient.style.display =
      ajoutclient.style.display === "flex" ? "none" : "flex";
  });
  closeBtn.addEventListener("click", function() {
    ajoutclient.style.display = "none";
  });
});

let currentStep = 1;

function nextStep() {
  document.getElementById("step" + currentStep).style.display = "none";
  currentStep++;
  document.getElementById("step" + currentStep).style.display = "block";
}

function prevStep() {
  document.getElementById("step" + currentStep).style.display = "none";
  currentStep--;
  document.getElementById("step" + currentStep).style.display = "block";
}

function enregistrer() {
  // Code pour enregistrer les données
  alert("Données enregistrées avec succès !");
}

function previewImage(event) {
  const preview = document.getElementById("preview");
  preview.style.display = "block";
  preview.src = URL.createObjectURL(event.target.files[0]);
}

/*****************************espace gestion des incidents */
document.addEventListener("DOMContentLoaded", function() {
  var interventionBTN = document.getElementById("interventionBTN");
  var intervention = document.getElementById("intervention");
  var closeBtninc = document.getElementById("closeBtninc");
  interventionBTN.addEventListener("click", function() {
    intervention.style.display =
      intervention.style.display === "flex" ? "none" : "flex";
  });
  closeBtninc.addEventListener("click", function() {
    intervention.style.display = "none";
  });
});

