// src/pages/ClientsList.jsx
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function ClientsList() {
  const [clients, setClients] = useState([]);
  const navigate = useNavigate();

  // Загружаем клиентов при монтировании
  useEffect(() => {
    fetch("/client") // эндпоинт из Go
      .then((res) => res.json())
      .then((data) => setClients(data))
      .catch((err) => console.error("Ошибка загрузки:", err));
  }, []);

  async function deleteClient(id) {
    if (!window.confirm("Удалить клиента?")) return;

    try {
      const res = await fetch(`/client/${id}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
      });

      if (res.ok) {
        alert("Клиент удален");
        setClients(clients.filter((c) => c.ID !== id));
      } else {
        const data = await res.json();
        alert(data.error || "Ошибка при удалении");
      }
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  }

  return (
    <div className="p-4">
      <h2 className="text-xl font-bold mb-4">Список клиентов</h2>

      <table className="border border-collapse w-full">
        <thead>
          <tr className="bg-gray-200">
            <th className="border p-2">Имя</th>
            <th className="border p-2">Фамилия</th>
            <th className="border p-2">Отчество</th>
            <th className="border p-2">Адрес</th>
            <th className="border p-2">Телефон</th>
            <th className="border p-2">Дата рождения</th>
            <th className="border p-2">Паспорт</th>
            <th className="border p-2">Действия</th>
          </tr>
        </thead>
        <tbody>
          {clients.map((client) => (
            <tr key={client.ID}>
              <td className="border p-2">{client.Firstname}</td>
              <td className="border p-2">{client.Lastname}</td>
              <td className="border p-2">{client.Middlename}</td>
              <td className="border p-2">{client.Address}</td>
              <td className="border p-2">{client.Phonenumber}</td>
              <td className="border p-2">
                {new Date(client.Dateofbirth).toLocaleDateString("ru-RU")}
              </td>
              <td className="border p-2">{client.Passport}</td>
              <td className="border p-2">
                <button
                  className="mr-2 text-blue-600"
                  onClick={() => navigate(`/client/edit/${client.ID}`)}
                >
                  ✏️ Редактировать
                </button>
                <button
                  className="text-red-600"
                  onClick={() => deleteClient(client.ID)}
                >
                  🗑️ Удалить
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      <br />
      <button
        className="bg-green-500 text-white px-3 py-1 rounded"
        onClick={() => navigate("/client/new")}
      >
        ➕ Добавить клиента
      </button>
    </div>
  );
}
