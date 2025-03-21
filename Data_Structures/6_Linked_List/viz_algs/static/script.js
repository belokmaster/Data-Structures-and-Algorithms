async function fetchList() {
    const response = await fetch("/list");
    const data = await response.json();
    renderList(data);
}

function renderList(node) {
    const container = document.getElementById("list");
    container.innerHTML = "";

    let current = node;
    while (current) {
        const div = document.createElement("div");
        div.classList.add("node");
        div.textContent = current.value;
        container.appendChild(div);

        if (current.next) {
            const arrow = document.createElement("span");
            arrow.textContent = " → ";
            container.appendChild(arrow);
        }

        current = current.next;
    }
}

async function addNode() {
    const value = document.getElementById("value").value;
    if (value === "") return;

    const response = await fetch("/list");
    const data = await response.json();
    
    // Если список не пуст, запускаем анимацию прохода
    if (data) await animateTraversal();

    // После анимации отправляем запрос на добавление
    await fetch("/add", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: `value=${encodeURIComponent(value)}`
    });

    // Обновляем список с задержкой для плавного появления нового элемента
    setTimeout(async () => {
        await fetchList();
        highlightLast();
    }, 200);
}

async function addNodeFront() {
    const value = document.getElementById("value").value;
    if (value === "") return;

    await fetch("/addFront", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: `value=${encodeURIComponent(value)}`
    });

    await fetchList();
    highlightFirst();
}

async function addNodeByIndex() {
    const value = document.getElementById("value").value;
    const index = document.getElementById("index").value;

    if (value === "" || index === "") return;

    // Анимация перед добавлением
    await animateTraversal();

    // Отправка запроса на добавление по индексу
    await fetch("/addIndex", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: `value=${encodeURIComponent(value)}&index=${encodeURIComponent(index)}`
    });

    // Обновляем список с задержкой
    setTimeout(async () => {
        await fetchList();
        highlightAtIndex(index); // Подсвечиваем добавленный элемент
    }, 200);
}

async function clearList() {
    await fetch("/clear", { method: "POST" });
    fetchList();
}

// Анимация прохода по списку (от начала к последнему элементу)
async function animateTraversal() {
    await fetchList(); // Обновляем список перед анимацией
    const nodes = document.querySelectorAll(".node");

    for (let i = 0; i < nodes.length; i++) {
        nodes[i].classList.add("highlight");
        await new Promise(resolve => setTimeout(resolve, 200));
        nodes[i].classList.remove("highlight");
    }
}

// Подсветка первого элемента (для добавления в начало)
function highlightFirst() {
    const nodes = document.querySelectorAll(".node");
    if (nodes.length > 0) {
        nodes[0].classList.add("highlight");
        setTimeout(() => nodes[0].classList.remove("highlight"), 300);
    }
}

// Подсветка последнего элемента (для добавления в конец)
function highlightLast() {
    const nodes = document.querySelectorAll(".node");
    if (nodes.length > 0) {
        const lastNode = nodes[nodes.length - 1];
        lastNode.classList.add("highlight");
        setTimeout(() => lastNode.classList.remove("highlight"), 300);
    }
}

// Подсветка добавленного элемента по индексу
function highlightAtIndex(index) {
    const nodes = document.querySelectorAll(".node");
    if (index < nodes.length) {
        const nodeToHighlight = nodes[index];
        nodeToHighlight.classList.add("highlight");
        setTimeout(() => nodeToHighlight.classList.remove("highlight"), 300);
    }
}

window.onload = fetchList;
