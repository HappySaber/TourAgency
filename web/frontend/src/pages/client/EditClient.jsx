import { useState } from "react";

export default function EditClientForm({ client, csrfToken }) {
  const [form, setForm] = useState({
    id: client?.id || "",
    firstname: client?.firstname || "",
    lastname: client?.lastname || "",
    middlename: client?.middlename || "",
    address: client?.address || "",
    phonenumber: client?.phonenumber || "",
    passport: client?.passport || "",
    dateofbirth: client?.dateofbirth
      ? client.dateofbirth.split("T")[0] // формат yyyy-MM-dd
      : "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const payload = {
      ...form,
      dateofbirth: new Date(form.dateofbirth).toISOString(),
    };

    try {
      const res = await fetch(`/client/edit/${form.id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": csrfToken,
        },
        body: JSON.stringify(payload),
      });

      if (res.ok) {
        alert("Клиент обновлен");
        window.location.href = "/client";
      } else {
        const data = await res.json();
        alert(data.error || "Ошибка при обновлении");
      }
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Редактирование клиента</h2>
      <form onSubmit={handleSubmit}>
        <input type="hidden" name="_csrf" value={csrfToken} />
        <input type="hidden" name="id" value={form.id} />

        <label>Имя:</label>
        <input
          type="text"
          name="firstname"
          value={form.firstname}
          onChange={handleChange}
          required
        />
        <br />

        <label>Фамилия:</label>
        <input
          type="text"
          name="lastname"
          value={form.lastname}
          onChange={handleChange}
          required
        />
        <br />

        <label>Отчество:</label>
        <input
          type="text"
          name="middlename"
          value={form.middlename}
          onChange={handleChange}
        />
        <br />

        <label>Адрес:</label>
        <input
          type="text"
          name="address"
          value={form.address}
          onChange={handleChange}
          required
        />
        <br />

        <label>Телефон:</label>
        <input
          type="tel"
          name="phonenumber"
          value={form.phonenumber}
          onChange={handleChange}
          pattern="\d{10}"
          placeholder="1234567890"
          required
        />
        <br />

        <label>Паспорт:</label>
        <input
          type="text"
          name="passport"
          value={form.passport}
          onChange={handleChange}
          required
        />
        <br />

        <label>Дата рождения:</label>
        <input
          type="date"
          name="dateofbirth"
          value={form.dateofbirth}
          onChange={handleChange}
          required
        />
        <br />

        <button type="submit">Сохранить изменения</button>
      </form>
    </div>
  );
}
