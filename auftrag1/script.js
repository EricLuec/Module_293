let input = document.querySelector("body");
input.addEventListener("change", calc);

function calc(e) {
    let startAnzahl = 100;
    let multiplicatorValue = document.getElementById("fname").value;

    document.getElementById("mehl").innerHTML = `${startAnzahl * multiplicatorValue}g Mehl`;
    document.getElementById("butter").innerHTML = `${startAnzahl * multiplicatorValue}g Butter`;
    document.getElementById("kaese").innerHTML = `${startAnzahl * multiplicatorValue}g Käse`;
    document.getElementById("personindicator").innerHTML = ` Dieses Rezept reicht für ${multiplicatorValue} Personen`;
}
    


