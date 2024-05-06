const StatusMessage = new URLSearchParams(window.location.search).get("r");
const myAlert = document.getElementById("loginForm");
const myAlert1 = document.getElementById("RegisterForm");
const overlay = document.createElement("div");
const EmailError = document.getElementById("EmailError");
const Message = document.getElementById("Message");


function showModal() {
  overlay.style.position = "fixed";
  overlay.style.zIndex = "10";
  overlay.style.left = "0%";
  overlay.style.top = "0%";
  overlay.style.width = "100%";
  overlay.style.height = "100%";
  overlay.style.backgroundColor = "rgba( 0, 0, 0, 0.5 )";

  document.body.appendChild(overlay);
  overlay.addEventListener("click", function () {
    myAlert.style.display = "none";
    myAlert1.style.display = "none";
    EmailError.textContent = "";
    this.remove(overlay);
  });
}

// Получаем элементы .card и <main>
var card = document.querySelector(".card");
var main = document.querySelector("main");

// Получаем размеры .card и <main>
var cardRect = card.getBoundingClientRect();
var mainRect = main.getBoundingClientRect();
var cardWidth = cardRect.width;
var mainHeight = mainRect.height;

// Функция для проверки расстояния между элементами
function isOverlap(newCircle, existingCircles) {
  for (var i = 0; i < existingCircles.length; i++) {
    var existingCircle = existingCircles[i];
    var rect1 = newCircle.getBoundingClientRect();
    var rect2 = existingCircle.getBoundingClientRect();
    var dx = rect1.left - rect2.left;
    var dy = rect1.top - rect2.top;
    var distance = Math.sqrt(dx * dx + dy * dy);
    if (distance < 100) {
      // Если расстояние меньше 50px, возвращаем true
      return true;
    }
  }
  return false;
}

var existingCircles = []; // Массив для хранения существующих кругов

// Создаем 5 элементов circle1 и добавляем их внутрь .card
for (var i = 0; i < 15; i++) {
  // Создаем новый элемент div
  var newCircle = document.createElement("div");

  // Добавляем класс "circle1" к созданному элементу
  newCircle.classList.add("circle1");

  // Генерируем случайные значения для размера, top и left
  var size = Math.floor(Math.random() * 120) + 1; // Размер от 1px до 150px
  var randomTop, randomLeft;

  // Перегенерируем координаты, если новый элемент пересекается с существующими
  do {
    randomTop = Math.floor(Math.random() * (mainHeight - 25)); // Отнимаем 25, чтобы элемент не выходил за границы
    randomLeft = Math.floor(Math.random() * (cardWidth - 25)); // Отнимаем 25, чтобы элемент не выходил за границы
    newCircle.style.top = randomTop + "px";
    newCircle.style.left = randomLeft + "px";
  } while (isOverlap(newCircle, existingCircles));

  // Добавляем созданный элемент внутрь .card
  card.appendChild(newCircle);

  // Добавляем созданный элемент в массив существующих элементов
  existingCircles.push(newCircle);

  // Устанавливаем случайные значения в качестве стилей размера для нового элемента
  newCircle.style.width = size + "px";
  newCircle.style.height = size + "px";
}

function login() {
  myAlert1.style.display = "none";
  myAlert.style.zIndex = "100";
  myAlert.style.display = "block";
  myAlert.style.position = "fixed";
  myAlert.style.left = "50%";
  myAlert.style.top = "50%";
  myAlert.style.transform = "translate( -50% , -50% )";
  myAlert.style.boxShadow = " 0 , 0 , 10px , rgba( 0, 0, 0, 0.5 ) ";

  showModal();
}

function register() {
  myAlert.style.display = "none";
  myAlert1.style.zIndex = "100";
  myAlert1.style.display = "block";
  myAlert1.style.position = "fixed";
  myAlert1.style.left = "50%";
  myAlert1.style.top = "50%";
  myAlert1.style.transform = "translate( -50% , -50% )";
  myAlert1.style.boxShadow = " 0 , 0 , 10px , rgba( 0, 0, 0, 0.5 ) ";
  showModal();
}

switch (StatusMessage) {
  case "user/signup/$2a$12$.HjibCBIzegvrQ/dwDRzkO8DzBVEnFU6PlGlqhtCYu4hLMCgF.F3G":
    register();
    EmailError.textContent = "Sorry, this email is already taken";
    break;

  case "user/login/$2a$12$B1q8dQJdtB/chK3XMVZmFeWsVKds1t0Oyw8KmMsyDdkVpxEEz.zmS":
    login();
    Message.textContent = "You are successfully signed up. Please log in";
    break;

  case "user/login/$2a$12$qSexwUFOrmwkXucsSGdZu.1d0.YCu3/4gaxfVwdrRdb0b4ZnjP2de":
    login();
    Message.textContent = "Wrong email or password";
    Message.style.color = "red";

    break;

  case "user/login/$2a$12$TV0unRiB6mg0U.4FhoznvuyD8mR5wG9v5SZWi8.T8lMY0bvc5CiCy":
    login();
    break;

  case "user/signup/$2a$12$0Oun5XW9qwWIj1sKSfktn.w2ldBuRh9Btbjz.i3Kgnp8SVt.KC12K":
    register();
    break;

}


document.addEventListener("DOMContentLoaded", function () {
  const passwordInput = document.getElementById("password");
  const confirmPasswordInput = document.getElementById("confirmPassword");
  const passwordMatchError = document.getElementById("PasswordError");
  const submitButton = document.querySelector(".submit");

  function checkPasswordMatch() {
    const password = passwordInput.value;
    const confirmPassword = confirmPasswordInput.value;

    // Validate if password contains only allowed characters
    const containsAllowedCharacters = /^[a-zA-Z0-9!\"#$%&'()*+,\-.\/:;<=>?@[\\\]^_`{}~]+$/.test(password);

    if (password.length === 0) {
      // If password is empty, show error message and disable submit button
      passwordMatchError.textContent = "Password is required";
      submitButton.disabled = true;
    } else if (!containsAllowedCharacters) {
      // If password contains disallowed characters, show error message and disable submit button
      passwordMatchError.textContent = "Password contains disallowed characters. Allowed characters are A-Z, a-z, 0-9, and special characters: !\"#$%&'()*+,-./:;<=>?@[\\]^_`{}~";
      submitButton.disabled = true;
    } else if (password.length < 8 || password.length > 15) {
      // If password length is not within the required range, show error message and disable submit button
      passwordMatchError.textContent = "Password must be between 8 and 15 characters";
      submitButton.disabled = true;
    } else if (confirmPassword.length > 0 && password !== confirmPassword) {
      // If passwords do not match, show error message and disable submit button
      passwordMatchError.textContent = "Passwords do not match";
      submitButton.disabled = true;
    } else {
      // If all conditions pass, clear error message and enable submit button
      passwordMatchError.textContent = "";
      submitButton.disabled = false;
    }
  }

  // Add event listeners
  passwordInput.addEventListener("input", checkPasswordMatch);
  confirmPasswordInput.addEventListener("input", checkPasswordMatch);
});

const inputField = document.getElementById('myInput');
const textareaField = document.getElementById('myTextarea');
const categoryCheckboxes = document.querySelectorAll('input[name="category"]');
const submitButton = document.querySelector('.submit2');
const categoryError = document.getElementById('category-error');

const MAX_TITLE_LENGTH = 100;
const MAX_DESCRIPTION_LENGTH = 1001;

function checkFields() {
  let titleFilled = inputField.value.trim().length > 0 && inputField.value.trim().length <= MAX_TITLE_LENGTH;
  let descriptionFilled = textareaField.value.trim().length > 0 && textareaField.value.trim().length <= MAX_DESCRIPTION_LENGTH;
  let categorySelected = false;

  // Check if at least one category is selected
  categoryCheckboxes.forEach(checkbox => {
    if (checkbox.checked) {
      categorySelected = true;
    }
  });

  console.log('Title Filled:', titleFilled);
  console.log('Description Filled:', descriptionFilled);
  console.log('Category Selected:', categorySelected);

  // Display error message only when both title and description are filled
  if (titleFilled && descriptionFilled && !categorySelected) {
    categoryError.style.display = 'block';
  } else {
    categoryError.style.display = 'none';
  }

  // Enable/disable submit button based on all fields being filled
  submitButton.disabled = !(titleFilled && descriptionFilled && categorySelected);
}

// Add event listeners to input fields and checkboxes
inputField.addEventListener('input', checkFields);
textareaField.addEventListener('input', checkFields);
categoryCheckboxes.forEach(checkbox => {
  checkbox.addEventListener('change', checkFields);
});

// Initially disable the submit button
checkFields();
