INSERT INTO articles (image_urls, title, uploaded_time, tags, content, created_at)
VALUES 
    ('{"big_img_url": "https://example.com/images/healthy-big1.jpg", "small_img_url": "https://example.com/images/healthy-small1.jpg"}', 
     'Healthy Eating for Beginners', 
     '2024-11-20T10:00:00Z', 
     'healthy, eating, beginners', 
     'This article provides an introduction to healthy eating and how to get started.', 
     '2024-11-20T10:00:00Z'),

    ('{"big_img_url": "https://example.com/images/fitness-big2.jpg", "small_img_url": "https://example.com/images/fitness-small2.jpg"}', 
     'The Power of Exercise for Mental Health', 
     '2024-11-19T09:30:00Z', 
     'exercise, mental health, wellness', 
     'This article explores the mental health benefits of regular exercise and how to incorporate it into daily life.', 
     '2024-11-19T09:30:00Z'),

    ('{"big_img_url": "https://example.com/images/stress-big3.jpg", "small_img_url": "https://example.com/images/stress-small3.jpg"}', 
     'How to Manage Stress Effectively', 
     '2024-11-18T11:00:00Z', 
     'stress, management, mindfulness', 
     'In this article, we discuss various stress management techniques, including mindfulness and time management practices.', 
     '2024-11-18T11:00:00Z'),

    ('{"big_img_url": "https://example.com/images/nutrition-big4.jpg", "small_img_url": "https://example.com/images/nutrition-small4.jpg"}', 
     'Nutrition Tips for a Balanced Diet', 
     '2024-11-17T14:45:00Z', 
     'nutrition, diet, health', 
     'This article covers essential nutrition tips to maintain a balanced and healthy diet, including food groups and portion sizes.', 
     '2024-11-17T14:45:00Z'),

    ('{"big_img_url": "https://example.com/images/mindfulness-big5.jpg", "small_img_url": "https://example.com/images/mindfulness-small5.jpg"}', 
     'Mindfulness Practices for Everyday Life', 
     '2024-11-16T08:30:00Z', 
     'mindfulness, meditation, stress relief', 
     'A guide to incorporating mindfulness practices into your daily routine to help reduce stress and increase mental clarity.', 
     '2024-11-16T08:30:00Z');

INSERT INTO global_config (key, value, created_at, scope)
VALUES 
    ('api-version', 'v0.0.9', strftime('%Y-%m-%dT%H:%M:%SZ', 'now'), 'global');