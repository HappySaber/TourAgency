import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const EditProvider = ({ providerData, csrfToken }) => {
  const [form, setForm] = useState({
    id: providerData.ID,
    name: providerData.Name,
    addressto: providerData.Addressto,
    address: providerData.Address,
    email: providerData.Email,
    phonenumber: providerData.PhoneNumber,
  });

  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch(`/provider/edit/${form.id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": csrfToken,
        },
        body: JSON.stringify(form),
      });

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.error || data.message || "Ошибка при обновлении");
      }

      alert(data.message || "Провайдер обновлен");
      navigate("/provider"); // редирект на список провайдеров
    } catch (err) {
      console.error("Ошибка при обновлении провайдера:", err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Редактировать провайдера</h2>
      <form onSubmit={handleSubmit}>
        <label>Название:</label><br />
        <input type="text" name="name" value={form.name} onChange={handleChange} required /><br /><br />

        <label>Адрес назначения:</label><br />
        <input type="text" name="addressto" value={form.addressto} onChange={handleChange} required /><br /><br />

        <label>Адрес:</label><br />
        <input type="text" name="address" value={form.address} onChange={handleChange} required /><br /><br />

        <label>Email:</label><br />
        <input type="email" name="email" value={form.email} onChange={handleChange} required /><br /><br />

        <label>Телефон:</label><br />
        <input type="text" name="phonenumber" value={form.phonenumber} onChange={handleChange} required /><br /><br />

        <button type="submit">Сохранить изменения</button>
      </form>
    </div>
  );
};

export default EditProvider;
