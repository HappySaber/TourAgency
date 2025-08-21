import React, { useState } from "react";

export default function CreateClient() {
  const [formData, setFormData] = useState({
    firstname: "",
    lastname: "",
    middlename: "",
    address: "",
    phonenumber: "",
    passport: "",
    dateofbirth: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const client = {
        ...formData,
        // если сервер ждёт ISO-дату — оставляем так
        dateofbirth: new Date(formData.dateofbirth).toISOString(),
      };

      const res = await fetch("/client/new", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify(client),
      });

      if (!res.ok) {
        const errorData = await res.json().catch(() => ({}));
        throw new Error(
          errorData.error || errorData.message || `HTTP error! ${res.status}`
        );
      }

      const data = await res.json();
      alert(data.message || "Клиент успешно создан");
      window.location.href = "/client";
    } catch (err) {
      console.error("Ошибка при создании клиента:", err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div className="p-4 max-w-xl mx-auto">
      <h2 className="text-2xl font-bold mb-4">Добавление клиента</h2>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label>Имя:</label>
          <input
            type="text"
            name="firstname"
            value={formData.firstname}
            onChange={handleChange}
            required
            className="border p-2 w-full rounded"
          />
        </div>

        <div>
          <label>Фамилия:</label>
          <input
            type="text"
            name="lastname"
            value={formData.lastname}
            onChange={handleChange}
            required
            className="border p-2 w-full rounded"
          />
        </div>

        <div>
          <label>Отчество:</label>
          <input
            type="text"
            name="middlename"
            value={formData.middlename}
            onChange={handleChange}
            className="border p-2 w-full rounded"
          />
        </div>

        <div>
          <label>Адрес:</label>
          <input
            type="text"
            name="address"
            value={formData.address}
            onChange={handleChange}
            required
            className="border p-2 w-full rounded"
          />
        </div>

        <div>
          <label>Телефон:</label>
          <input
            type="tel"
            name="phonenumber"
            value={formData.phonenumber}
            onChange={handleChange}
            pattern="[0-9]{10}"
            maxLength="10"
            required
            title="10 цифр номера без пробелов и спецсимволов"
            className="border p-2 w-full rounded"
          />
        </div>

        <div>
          <label>Паспорт:</label>
          <input
            type="text"
            name="passport"
            value={formData.passport}
            onChange={handleChange}
            required
            className="border p-2 w-full rounded"
          />
        </div>

        <div>
          <label>Дата рождения:</label>
          <input
            type="date"
            name="dateofbirth"
            value={formData.dateofbirth}
            onChange={handleChange}
            required
            className="border p-2 w-full rounded"
          />
        </div>

        <button
          type="submit"
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        >
          Сохранить
        </button>
      </form>
    </div>
  );
}
