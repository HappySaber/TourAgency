import React, { useState } from "react";

export default function AddToursForm({ consultationId, tours = [] }) {
  const [selectedTours, setSelectedTours] = useState(
    tours.map(t => ({
      ...t,
      checked: t.Checked || false,
      discount: t.Discount || "",
      quantity: t.Quantity || ""
    }))
  );

  const handleCheckboxChange = (id) => {
    setSelectedTours(prev =>
      prev.map(t =>
        t.ID === id ? { ...t, checked: !t.checked } : t
      )
    );
  };

  const handleInputChange = (id, field, value) => {
    setSelectedTours(prev =>
      prev.map(t =>
        t.ID === id ? { ...t, [field]: value } : t
      )
    );
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const formData = new URLSearchParams();
    selectedTours.forEach(t => {
      if (t.checked) {
        formData.append("tours", t.ID);
        formData.append("discount_" + t.ID, t.discount);
        formData.append("quantity_" + t.ID, t.quantity);
      }
    });

    try {
      const res = await fetch(`/consultation/${consultationId}/tours`, {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
        body: formData.toString()
      });

      if (res.ok) {
        alert("Туры успешно добавлены");
        window.location.href = "/consultation";
      } else {
        const data = await res.json();
        alert(data.error || "Ошибка при добавлении туров");
      }
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Добавление туров к консультации</h2>
      <form onSubmit={handleSubmit}>
        <input type="hidden" name="consultation_id" value={consultationId} />

        {selectedTours.length > 0 ? (
          selectedTours.map(t => (
            <div key={t.ID}>
              <label>
                <input
                  type="checkbox"
                  checked={t.checked}
                  onChange={() => handleCheckboxChange(t.ID)}
                />
                {t.Name} — {t.Country}
              </label>
              {t.checked && (
                <div style={{ marginLeft: 20 }}>
                  Скидка:{" "}
                  <input
                    type="text"
                    value={t.discount}
                    onChange={(e) =>
                      handleInputChange(t.ID, "discount", e.target.value)
                    }
                  />
                  <br />
                  Кол-во:{" "}
                  <input
                    type="text"
                    value={t.quantity}
                    onChange={(e) =>
                      handleInputChange(t.ID, "quantity", e.target.value)
                    }
                  />
                </div>
              )}
            </div>
          ))
        ) : (
          <p>Нет доступных туров.</p>
        )}

        <br />
        <button type="submit">Сохранить туры</button>
      </form>
    </div>
  );
}
