-- create resume

INSERT INTO 
resumes
(
    id,
    user_id,
    job_title,
    created_at,
    updated_at,
    version
)
values(
    1,
    1,
    "Web Developer",
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
);

-- create reference list

insert into
    reference_lists
    (
        id,
        resume_id,
        created_at,
        updated_at,
        version
    )
    values(
        1,
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP,
        1
    );

-- SQL Migration: Insert fake reference items

INSERT INTO reference_items (
    ref_list_id,
    first_name,
    last_name,
    email,
    phone_number,
    job_title,
    organization,
    type,
    address_1,
    address_2,
    city,
    state,
    zipcode,
    content,
    created_at,
    updated_at,
    version
) VALUES
(
    1, -- Replace with appropriate ref_list_id
    'John',
    'Doe',
    'john.doe@example.com',
    '555-1234',
    'Software Engineer',
    'Tech Corp Inc.',
    'professional',
    '123 Main St',
    'Apt 4B',
    'Anytown',
    'CA',
    '90210',
    'John is a highly skilled software engineer with five years of experience in full-stack development.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Replace with appropriate ref_list_id
    'Jane',
    'Smith',
    'jane.smith@college.edu',
    '555-5678',
    'Lecturer',
    'University of Learning',
    'academic',
    '456 Elm St.',
    '',
    'Metropolis',
    'NY',
    '10101',
    'Jane has been a prominent lecturer in computer science for the past decade, known for her innovative teaching methods.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Replace with appropriate ref_list_id
    'Alice',
    'Brown',
    'alice.brown@bizexample.com',
    '555-8765',
    'Project Manager',
    'Business Solutions Ltd.',
    'industry',
    '789 Maple Ave',
    '',
    'Big City',
    'TX',
    '75001',
    'Alice is a proficient project manager, leading diverse teams to success in the IT sector for over ten years.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
);

insert into
    award_lists
    (
    id,
    resume_id,
    created_at,
    updated_at,
    version
    )
    values(
    1,
    1,
    current_timestamp,
    current_timestamp,
    1
    );


-- award list items

-- SQL Migration: Insert fake award items

INSERT INTO award_items (
    award_list_id,
    title,
    org_name,
    received_year,
    content,
    created_at,
    updated_at,
    version
) VALUES
(
    1, -- Replace with the appropriate award_list_id
    'Outstanding Developer Award',
    'Tech Excellence Institute',
    2023,
    'This prestigious award was given in recognition of exceptional software development achievements and contributions.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Replace with the appropriate award_list_id
    'Innovative Educator of the Year',
    'National Teaching Awards Association',
    2022,
    'Awarded for pioneering educational techniques that have significantly advanced student engagement and outcomes in computer science.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Replace with the appropriate award_list_id
    'Leadership in Project Management',
    'International Management Professionals',
    2024,
    'Recognized for outstanding leadership and effectiveness in delivering high-impact projects across various sectors.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
);


-- education lists
insert into
education_lists
(
resume_id,
created_at,
updated_at,
version
)
values(
1,
current_timestamp,
current_timestamp,
1
);


-- education list items

-- SQL Migration: Insert dummy education items

INSERT INTO education_items (
    education_lists_id,
    name,
    degree_year,
    degree,
    address_1,
    address_2,
    city,
    state,
    zipcode,
    created_at,
    updated_at,
    version
) VALUES
(
    1, -- Assumed education_lists_id for these entries
    'Massachusetts Institute of Technology',
    2020,
    'Bachelor of Science in Computer Science',
    '77 Massachusetts Ave',
    '',
    'Cambridge',
    'MA',
    '02139',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Assumed education_lists_id for these entries
    'Stanford University',
    2022,
    'Master of Science in Artificial Intelligence',
    '450 Serra Mall',
    '',
    'Stanford',
    'CA',
    '94305',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Assumed education_lists_id for these entries
    'University of Oxford',
    2018,
    'Bachelor of Arts in Philosophy, Politics and Economics',
    'Wellington Square',
    '',
    'Oxford',
    'Oxfordshire',
    'OX1 2JD',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
);


-- employment list

insert into 
    employment_lists
(
    resume_id,
    created_at,
    updated_at,
    version
)
values(
1,
current_timestamp,
current_timestamp,
1
);

-- SQL Migration: Insert dummy employment list items

INSERT INTO employment_list_items (
    employment_lists_id,
    title,
    date_from,
    date_to,
    job_title,
    summary,
    created_at,
    updated_at,
    version
) VALUES
(
    1, -- Assumed employment_lists_id for these entries
    'Software Engineer at Example Corp',
    2018,
    2021,
    'Software Engineer',
    'Developed and maintained web applications, leading a team of junior developers in adopting new technologies to improve software efficiency.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Assumed employment_lists_id for these entries
    'Project Manager at Global Solutions',
    2021,
    2023,
    'Project Manager',
    'Oversaw multiple project deliveries, streamlined operations across departments, and improved client satisfaction through effective resource management.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Assumed employment_lists_id for these entries
    'Data Analyst at Data Insights Ltd',
    2015,
    2018,
    'Data Analyst',
    'Analyzed large datasets to identify trends, created detailed reports that improved business decision-making, and collaborated with cross-functional teams to implement data-driven strategies.',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
);

-- skill list

insert into
skill_lists
(
resume_id,
title,
created_at,
updated_at,
version
)
values(
    1,
    "Skill List",
    current_timestamp,
    current_timestamp,
    1
);

-- SQL Migration: Insert dummy skill list items

INSERT INTO skill_list_items (
    skill_lists_id,
    title,
    content,
    duration,
    created_at,
    updated_at,
    version
) VALUES
(
    1, -- Assumed skill_lists_id for these entries
    'Advanced Python Programming',
    'Expertise in writing complex Python scripts for data analysis, automation, and web development, with five years of practical experience.',
    2019,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Assumed skill_lists_id for these entries
    'Project Management',
    'Proficient in managing large-scale projects, utilizing methodologies such as Agile and Scrum to enhance team productivity and project outcomes.',
    2020,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Assumed skill_lists_id for these entries
    'Machine Learning',
    'Skilled in designing and implementing machine learning models, achieving operational improvements and predictive analytics insights.',
    2021,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
);

-- insert an objective

insert into 
    objectives
    (
        resume_id,
        content,
        created_at,
        updated_at,
        version
    )
    values(
        1,
        "Some long winded content about how I want to get a programming job that aligns with my personal value system",
        current_timestamp,
        current_timestamp,
        1
    );

-- social media list

insert into
social_media_lists
(
resume_id,
created_at,
updated_at,
version
)
values(
    1,
    current_timestamp,
    current_timestamp,
    1

);

-- social media items

-- SQL Migration: Insert dummy social media list items

INSERT INTO social_media_list_items (
    social_media_lists_id,
    company_name,
    username,
    web_address,
    created_at,
    updated_at,
    version
) VALUES
(
    1, -- Assumed social_media_lists_id for these entries
    'Twitter',
    'JohnDoeTech',
    'https://twitter.com/JohnDoeTech',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Assumed social_media_lists_id for these entries
    'LinkedIn',
    'jane-smith-lecturer',
    'https://www.linkedin.com/in/jane-smith-lecturer',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
),
(
    1, -- Assumed social_media_lists_id for these entries
    'GitHub',
    'alicebrown-dev',
    'https://github.com/alicebrown-dev',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
);

-- SQL Migration: Insert a dummy resume contact detail

INSERT INTO resume_contact_details (
    resume_id,
    first_name,
    last_name,
    address_1,
    address_2,
    city,
    state,
    zipcode,
    email,
    phone_number,
    web_address,
    created_at,
    updated_at,
    version
) VALUES
(
    1, -- Assumed resume_id for this entry
    'Alex',
    'Johnson',
    '123 Elm Street',
    'Apt 5',
    'Springfield',
    'IL',
    '62704',
    'alex.johnson@example.com',
    '555-9876',
    'https://alexjohnson.dev',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    1
);









