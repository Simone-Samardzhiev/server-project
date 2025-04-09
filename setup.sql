CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    email    VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)        NOT NULL
);

CREATE TABLE events
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    date        DATE         NOT NULL,
    address     VARCHAR(255) NOT NULL,
    image_url   VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL
);

INSERT INTO events(title, date, address, image_url, description)
VALUES ('Доброволческа акция: Почисти природата',
        '2025-04-15',
        'Парк „Борисова градина“, София',
        'https://i.postimg.cc/4xptQzXm/image1.avif',
        'Присъединете се към нас в еднодневна инициатива за почистване на парка! Ще осигурим ръкавици и чували, а в края на деня ще има пикник за участниците.'),
       ('Как да създаваме екологични навици?',
        '2025-04-25',
        'Зала „Грийн Лаб“, Пловдив',
        'https://i.postimg.cc/TYVgrVdh/image2.avif',
        'Практически семинар за изграждане на устойчиви навици в ежедневието – от намаляване на отпадъците до екологично пазаруване. Включени са интерактивни задачи и подаръци за участниците.'),
       ('Фестивал на уличното изкуство',
        '2025-05-05',
        'Център на Варна',
        'https://i.postimg.cc/dtPGkHKz/image3.avif',
        ' Потопи се в света на уличното изкуство! Очакват те графити демонстрации, музика на живо и арт базар с местни творци.');

CREATE TABLE registrations
(
    id       SERIAL PRIMARY KEY,
    user_id  INT NOT NULL REFERENCES users (id),
    event_id INT NOT NULL REFERENCES events (id)
);

CREATE TABLE gallery
(
    id        SERIAL PRIMARY KEY,
    image_url VARCHAR(255) NOT NULL
);

INSERT INTO gallery (image_url)
VALUES ('https://i.postimg.cc/L5djggxj/temp-Image135-Xk-N.avif'),
       ('https://i.postimg.cc/K8Y3z5md/temp-Image3-So-YMI.avif'),
       ('https://i.postimg.cc/HxMymsTn/temp-Image4tqt9j.avif'),
       ('https://i.postimg.cc/prtnLbdZ/temp-Image505oa-K.avif'),
       ('https://i.postimg.cc/XvJfRbZB/temp-Image5g-N24i.avif'),
       ('https://i.postimg.cc/TwQXpnd4/temp-Image6-AUJy-B.avif'),
       ('https://i.postimg.cc/8C0RD0WK/temp-Image6yu0-Xe.avif'),
       ('https://i.postimg.cc/h4xLWMGs/temp-Image7-Z68-PK.avif'),
       ('https://i.postimg.cc/28rZpDMj/temp-Image7z-Inpb.avif'),
       ('https://i.postimg.cc/QMTt94wn/temp-Image9-Pgenl.avif'),
       ('https://i.postimg.cc/Y9xtYJ38/temp-Image-B1n7-Dj.avif'),
       ('https://i.postimg.cc/DzHXRbkR/temp-Imageb40-On-L.avif'),
       ('https://i.postimg.cc/Gp8BPh8K/temp-Image-BA9z-Mi.avif'),
       ('https://i.postimg.cc/WzbTnXYV/temp-Imageb-UHUz-E.avif'),
       ('https://i.postimg.cc/jdSSQdBf/temp-Imageccz-SOS.avif'),
       ('https://i.postimg.cc/NM8jbyg6/temp-Imaged-XVP5z.avif'),
       ('https://i.postimg.cc/SKCSNXYB/temp-Imageem-Pij-O.avif'),
       ('https://i.postimg.cc/1XJ5hLV4/temp-Imagegar7-Bl.avif'),
       ('https://i.postimg.cc/2j2z6N3S/temp-Imagegy0asp.avif'),
       ('https://i.postimg.cc/nhMMZckx/temp-Image-GYs-Jjn.avif'),
       ('https://i.postimg.cc/XYxK1Jv0/temp-Image-H1u-Ky-B.avif'),
       ('https://i.postimg.cc/GhX9Zsvz/temp-Imageh-Kwg-Hz.avif'),
       ('https://i.postimg.cc/tC3NPDQJ/temp-Image-Ijd-DU4.avif'),
       ('https://i.postimg.cc/1XcVgCnZ/temp-Image-IL71-EI.avif'),
       ('https://i.postimg.cc/pdpj0J7D/temp-Image-Knszsx.avif'),
       ('https://i.postimg.cc/C1njKssF/temp-Imagek-PDHQc.avif'),
       ('https://i.postimg.cc/8CQG31Xt/temp-Imageksnh2-M.avif'),
       ('https://i.postimg.cc/kXmxn86m/temp-Image-L4-Bw-OF.avif'),
       ('https://i.postimg.cc/h4rGd7Zg/temp-Imagemz-Ib-R5.avif'),
       ('https://i.postimg.cc/YCV0SKnW/temp-Image-N0r-Dmc.avif'),
       ('https://i.postimg.cc/bN5px0xX/temp-Image-N7p4n5.avif'),
       ('https://i.postimg.cc/Y9tHrdHF/temp-Imagen-Dc-NIr.avif'),
       ('https://i.postimg.cc/2y6Nv338/temp-Image-O9wb-Ny.avif'),
       ('https://i.postimg.cc/zBNJLJnV/temp-Image-OB6e0p.avif'),
       ('https://i.postimg.cc/fR6WN7y7/temp-Imageo-POQhv.avif'),
       ('https://i.postimg.cc/9X8h2PN4/temp-Image-OQK4-Lu.avif'),
       ('https://i.postimg.cc/fyrnvZ07/temp-Imagep0ris-H.avif'),
       ('https://i.postimg.cc/1RpqYtWY/temp-Image-PK1-FQA.avif'),
       ('https://i.postimg.cc/kg7Fy8qc/temp-Imagepudp-B1.avif'),
       ('https://i.postimg.cc/VkHrbb65/temp-Image-Qli-XZ8.avif'),
       ('https://i.postimg.cc/jjHTYGGK/temp-Imageqwu-CJs.avif'),
       ('https://i.postimg.cc/rwhWCPhV/temp-Image-RKX4-Ic.avif'),
       ('https://i.postimg.cc/Vksnt8WZ/temp-Imager-NYqn-L.avif'),
       ('https://i.postimg.cc/qMx2gRyC/temp-Images5io-By.avif'),
       ('https://i.postimg.cc/TwnpS7cW/temp-Images-At1j-U.avif'),
       ('https://i.postimg.cc/4yJN0Jdr/tempImagetTn41o.avif'),
       ('https://i.postimg.cc/4xhZBKPY/tempImageuJNknV.avif'),
       ('https://i.postimg.cc/DypyC3KM/tempImageurUSwS.avif'),
       ('https://i.postimg.cc/t41RwvqC/tempImageuuc18z.avif'),
       ('https://i.postimg.cc/x8KMvjZh/tempImageVYEn7a.avif'),
       ('https://i.postimg.cc/NF4LGKSs/tempImagexDRpYv.avif'),
       ('https://i.postimg.cc/9f6DL5yh/tempImagexswQy3.avif'),
       ('https://i.postimg.cc/Y0nkyYTW/tempImageYMhZCR.avif'),
       ('https://i.postimg.cc/Dy074pbd/tempImageypqvzw.avif'),
       ('https://i.postimg.cc/SRqJp6Z1/tempImageYWGKbD.avif'),
       ('https://i.postimg.cc/rFjtMH21/tempImageZQUKPp.avif'),
       ('https://i.postimg.cc/Y9BS1Nw2/tempImageZUSsNX.avif');






