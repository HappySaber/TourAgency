import React, { useState } from "react";

const AddServicesForm = ({ consultationId, allServices = [] }) => {
  const [services, setServices] = useState(
    allServices.map((s) => ({
      id: s.ID,
      name: s.Name,
      checked: s.Checked || false,
      discount: s.Discount || "",
      quantity: s.Quantity || "",
    }))
  );

  const handleCheckboxChange = (id) => {
    setServices((prev) =>
      prev.map((s) =>
        s.id === id ? { ...s, checked: !s.checked } : s
      )
    );
  };

  const handleInputChange = (id, field, value) => {
    setServices((prev) =>
      prev.map((s) =>
        s.id === id ? { ...s, [field]: value } : s
      )
    );
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const selectedServices = services
      .filter((s) => s.checked)
      .map((s) => ({
        id: s.id,
        discount: s.discount,
        quantity: s.quantity,
      }));

    try {
      const res = await fetch(`/consultation/${consultationId}/services`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ consultation_id: consultationId, services: selectedServices }),
      });

      if (res.ok) {
        alert("Услуги сохранены");
        window.location.href = `/consultation/${consultationId}`;
      } else {
        const data = await res.json();
        alert(data.error || "Ошибка при сохранении");
      }
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Добавление услуг к консультации</h2>
      <form onSubmit={handleSubmit}>
        <input type="hidden" name="consultation_id" value={consultationId} />

        {services.length > 0 ? (
          services.map((s) => (
            <div key={s.id}>
              <label>
                <input
                  type="checkbox"
                  checked={s.checked}
                  onChange={() => handleCheckboxChange(s.id)}
                />
                {s.name}
              </label>
              {s.checked && (
                <div style={{ marginLeft: "20px" }}>
                  Скидка:{" "}
                  <input
                    type="text"
                    value={s.discount}
                    onChange={(e) =>
                      handleInputChange(s.id, "discount", e.target.value)
                    }
                  />
                  <br />
                  Кол-во:{" "}
                  <input
                    type="text"
                    value={s.quantity}
                    onChange={(e) =>
                      handleInputChange(s.id, "quantity", e.target.value)
                    }
                  />
                </div>
              )}
            </div>
          ))
        ) : (
          <p>Нет доступных услуг.</p>
        )}

        <br />
        <button type="submit">Сохранить услуги</button>
      </form>
    </div>
  );
};

export default AddServicesForm;
