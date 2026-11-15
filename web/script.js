const statusEl = document.getElementById("ws-status");
const ordersBody = document.getElementById("orders-body");
const orderForm = document.getElementById("order-form");

const URL = 'localhost:7192'
const API = 'http://localhost:7192'
const ws = new WebSocket(`ws://${URL}/ws`);

ws.onopen = () => {
  statusEl.textContent = "Connected";
  statusEl.className = "online";
};

ws.onclose = () => {
  statusEl.textContent = "Disconnected";
  statusEl.className = "offline";
};

ws.onmessage = (event) => {
  const order = JSON.parse(event.data);
  console.log("Raw JSON string:", event.data);
  console.log("Parsed JS object:", order);

  const productList = (order.product || [])
    .map((p) => `${p.name} (R ${p.price.toFixed(2)})`)
    .join(", ");

  const unitTotal = (order.product || []).reduce((sum, p) => sum + p.price, 0);
  const totalPrice = unitTotal * order.quantity;

  const row = document.createElement("tr");
  row.innerHTML = `
    <td title="${order.order_id}">${order.order_id.slice(0, 8)}...</td>
    <td>${productList}</td>
    <td>${order.quantity}</td>
    <td>R ${totalPrice.toFixed(2)}</td>
    <td><span class="badge-pending">${order.status}</span></td>
    <td>${order.purchased_at}</td>
  `;
  ordersBody.prepend(row);
};

orderForm.addEventListener("submit", async (e) => {
  e.preventDefault();

  const name = document.getElementById("name").value;
  const price = parseFloat(document.getElementById("price").value);
  const quantity = parseInt(document.getElementById("quantity").value, 10);

  const payload = {
    product: [{ name, price }],
    quantity: quantity,
  };

  await fetch(`${API}/api/orders`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  orderForm.reset();
});