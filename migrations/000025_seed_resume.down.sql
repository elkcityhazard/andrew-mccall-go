-- Delete from resume_contact_details
DELETE FROM resume_contact_details WHERE resume_id = 1 AND first_name = 'Alex' AND last_name = 'Johnson';

-- Delete from social_media_list_items
DELETE FROM social_media_list_items WHERE social_media_lists_id = 1;

-- Delete from social_media_lists
DELETE FROM social_media_lists WHERE resume_id = 1;

-- Delete from objectives
DELETE FROM objectives WHERE resume_id = 1;

-- Delete from skill_list_items
DELETE FROM skill_list_items WHERE skill_lists_id = 1;

-- Delete from skill_lists
DELETE FROM skill_lists WHERE resume_id = 1 AND title = 'Skill List';

-- Delete from employment_list_items
DELETE FROM employment_list_items WHERE employment_lists_id = 1;

-- Delete from employment_lists
DELETE FROM employment_lists WHERE resume_id = 1;

-- Delete from education_items
DELETE FROM education_items WHERE education_lists_id = 1;

-- Delete from education_lists
DELETE FROM education_lists WHERE resume_id = 1;

-- Delete from award_items
DELETE FROM award_items WHERE award_list_id = 1;

-- Delete from award_lists
DELETE FROM award_lists WHERE resume_id = 1;

-- Delete from reference_items
DELETE FROM reference_items WHERE ref_list_id = 1;

-- Delete from reference_lists
DELETE FROM reference_lists WHERE resume_id = 1;

-- Delete from resumes
DELETE FROM resumes WHERE id = 1 AND user_id = 1;

