IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'b12a4ead-9571-4121-92d4-ba94228cb3e4') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'b12a4ead-9571-4121-92d4-ba94228cb3e4', N'Evacuation', N'Emergency medical evacuation ', 1, 0, 4, NULL, N'5263c566-9a44-41bc-a1ca-f1d9e69c49f9', N'https://gobear0sg0prod.blob.core.windows.net/coveragelogos/emergency_medical_evacuation_and_repatriation.svg?rnd=1418958373275', 0, NULL, N'0,20000,100000,250000,500000,1000000,~', NULL, 0, 0, N'All', 1, 0, 3)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'cfd3adec-54c7-4b0f-ac08-fd1a72d6b8fb') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'cfd3adec-54c7-4b0f-ac08-fd1a72d6b8fb', N'LossValuables', N'Loss or damage to belongings', 1, 0, 6, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', N'https://gobear0sg0prod.blob.core.windows.net/coveragelogos/loss_of_valuables,_baggage_and_it_device.svg?rnd=1418960334860', 0, NULL, N'0,1000,3000,5000,8000', NULL, 0, 0, N'All', 1, 0, 4)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '23b4cfeb-f799-4dfa-a720-69b7c9dc6d49') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'23b4cfeb-f799-4dfa-a720-69b7c9dc6d49', N'Repatriationofmortalremains', N'Send your remains home', 0, 0, 1, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '21313a4c-4d12-4e41-b240-d9a45006b706') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'21313a4c-4d12-4e41-b240-d9a45006b706', N'ME(S)', N'Follow-up treatment back home ', 0, 0, 4, NULL, N'2364146f-ebb7-4ca3-9e70-255fc0c8c88c', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '1c2b573b-a021-446b-bccb-470c91f2f8bb') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'1c2b573b-a021-446b-bccb-470c91f2f8bb', N'LeisureSports', N'Leisure sports coverage', 0, 0, 2, NULL, N'2364146f-ebb7-4ca3-9e70-255fc0c8c88c', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'de5c0ebd-dbb1-4295-aee8-701ba4c81efb') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'de5c0ebd-dbb1-4295-aee8-701ba4c81efb', N'CancelledTrip', N'Cancelled Trip', 0, 0, 2, NULL, N'dacf99d2-be0b-41ee-80e0-8ff633cfd607', NULL, 0, NULL, N'0,2000,5000,10000,15000,20000', NULL, 1, 0, N'All', 1, 0, 3)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '3078798f-4a33-474d-aa69-d4973ecac850') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'3078798f-4a33-474d-aa69-d4973ecac850', N'BJ', N'Bungee jumping', 0, 0, 3, NULL, N'2179424b-36d0-413c-89f6-d2102916757c', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '925dd95b-445e-4d8c-8aa4-129f74d74644') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'925dd95b-445e-4d8c-8aa4-129f74d74644', N'Missed Connection', N'Missed Connection', 0, 0, 10, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', NULL, 0, NULL, N'0,1000', NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '3ade06a6-9207-45c3-b166-12b55055c879') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'3ade06a6-9207-45c3-b166-12b55055c879', N'MedicalOverseas', N'Overseas Medical Expenses', 1, 0, 6, NULL, N'2364146f-ebb7-4ca3-9e70-255fc0c8c88c', N'https://gobear0sg0prod.blob.core.windows.net/coveragelogos/medical_expense_incurred_overseas.svg?rnd=1418960236793', 0, NULL, N'0,10000,100000,200000,500000,1000000,2000000', NULL, 0, 0, N'Individual', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'acd1436c-d69d-4cd7-9900-9a5917845d41') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'acd1436c-d69d-4cd7-9900-9a5917845d41', N'LossTravelDocuments', N'Lost passport and other important docs', 0, 0, 5, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'a872ef60-26c7-4f90-9420-c353a7b50d25') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'a872ef60-26c7-4f90-9420-c353a7b50d25', N'SDiving', N'Skydiving', 0, 0, 4, NULL, N'2179424b-36d0-413c-89f6-d2102916757c', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '1c4b7f5b-84c4-4439-afc1-1571651cb4d9') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'1c4b7f5b-84c4-4439-afc1-1571651cb4d9', N'DelayedTrip', N'Delayed Trip', 0, 0, 1, NULL, N'dacf99d2-be0b-41ee-80e0-8ff633cfd607', NULL, 0, NULL, N'0,200,600,1000,2000,3000', NULL, 1, 0, N'All', 1, 0, 4)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '77dcf9fc-02f3-4ee2-afa1-cfff24f1b8fe') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'77dcf9fc-02f3-4ee2-afa1-cfff24f1b8fe', N'Trip Shortening', N'Interrupted trip', 0, 0, 12, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '1692b21b-7a1f-403b-b4f0-1095184fabe6') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'1692b21b-7a1f-403b-b4f0-1095184fabe6', N'Skiing', N'Skiing (excludes off piste)', 0, 0, 2, NULL, N'2179424b-36d0-413c-89f6-d2102916757c', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 0, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '226ef9cc-137f-41b5-ae27-2212df642aec') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'226ef9cc-137f-41b5-ae27-2212df642aec', N'Sky Diving', N'Sky diving', 0, 0, 4, NULL, N'60b91b43-1f6f-4468-9b9e-7b0eba0950a0', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '225ecc18-efea-42ea-87d8-644b3829e508') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'225ecc18-efea-42ea-87d8-644b3829e508', N'DeathDisablement', N'Personal accident payment', 1, 0, 7, NULL, N'2364146f-ebb7-4ca3-9e70-255fc0c8c88c', N'https://gobear0sg0prod.blob.core.windows.net/coveragelogos/personal_accident,_death_and_permanent_disablement.svg?rnd=1418960253191', 0, NULL, N'0,50000,100000,200000,300000,500000', NULL, 0, 0, N'All', 1, 0, 1)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '3b4bf41b-1ff6-4d24-8438-b02d93cedaf2') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'3b4bf41b-1ff6-4d24-8438-b02d93cedaf2', N'Scuba Diving', N'Scuba diving', 0, 0, 5, NULL, N'60b91b43-1f6f-4468-9b9e-7b0eba0950a0', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'b2a3a1c3-a4e2-499e-b9cd-64555f3d1a98') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'b2a3a1c3-a4e2-499e-b9cd-64555f3d1a98', N'LossOrDamageToBelongings', N'Loss or Damage to Belongings', 0, 0, 4, NULL, N'dacf99d2-be0b-41ee-80e0-8ff633cfd607', NULL, 0, NULL, N'0,1000,3000,5000,8000', NULL, 1, 0, N'All', 1, 0, 1)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '7fc14c33-9830-4e3f-918d-7917d79f9b48') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'7fc14c33-9830-4e3f-918d-7917d79f9b48', N'HoT', N'Hiking or Trekking', 0, 0, 1, NULL, N'2179424b-36d0-413c-89f6-d2102916757c', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 0, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '4d796b8d-0dd3-403d-96c3-896e56a72b3c') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'4d796b8d-0dd3-403d-96c3-896e56a72b3c', N'OverseasBaggageDelay', N'Delayed baggage', 0, 0, 7, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'a821e4dc-706d-4dfe-b672-92c2e3e21c04') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'a821e4dc-706d-4dfe-b672-92c2e3e21c04', N'TravelDelay6Hours', N'Delayed trip', 0, 0, 15, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '137cff9c-db3b-4c25-ad39-f6477cbd5131') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'137cff9c-db3b-4c25-ad39-f6477cbd5131', N'Terrorismcoverage', N'Terrorism coverage', 0, 0, 1, NULL, N'2364146f-ebb7-4ca3-9e70-255fc0c8c88c', NULL, 0, NULL, N'0,10000,100000,200000,500000,1000000,2000000', NULL, 0, 0, N'Individual', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '34ee0ea0-bf44-451e-8b61-26514dcc3047') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'34ee0ea0-bf44-451e-8b61-26514dcc3047', N'PersonalLiability', N'If someone makes a claim against you', 0, 0, 5, NULL, N'5263c566-9a44-41bc-a1ca-f1d9e69c49f9', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '75bd15b2-fbdf-489f-892c-32b88b1611c6') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'75bd15b2-fbdf-489f-892c-32b88b1611c6', N'Skiing ', N'Skiing (excludes off piste)', 0, 0, 2, NULL, N'60b91b43-1f6f-4468-9b9e-7b0eba0950a0', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '24d9617c-c10a-4ba8-901a-09f23e52e944') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'24d9617c-c10a-4ba8-901a-09f23e52e944', N'RentalExcess', N'Rental car excess', 0, 0, 3, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'fa17f027-1328-4005-ac99-fdf79731cefe') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'fa17f027-1328-4005-ac99-fdf79731cefe', N'Delayed Baggage', N'Delayed Baggage', 0, 0, 3, NULL, N'dacf99d2-be0b-41ee-80e0-8ff633cfd607', NULL, 0, NULL, N'0,100,300,500,600,800', NULL, 1, 0, N'All', 1, 0, 2)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '25a9a0d8-10b9-4fd1-a53d-1c1f7c224b7e') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'25a9a0d8-10b9-4fd1-a53d-1c1f7c224b7e', N'Trip Postponement ', N'Postponed trip', 0, 0, 13, NULL, N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', NULL, 0, NULL, NULL, NULL, 0, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = '5fb3c05e-b8e3-4003-bc24-c55e97e78745') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'5fb3c05e-b8e3-4003-bc24-c55e97e78745', N'OverseasMedicalExpenses', N'Overseas Medical Expenses', 0, 0, 5, NULL, N'dacf99d2-be0b-41ee-80e0-8ff633cfd607', NULL, 0, NULL, N'0,10000,100000,200000,500000,1000000,2000000', NULL, 1, 0, N'All', 1, 0, 0)

IF NOT EXISTS(select 1 from dbo.CoverageItemDefinitions where ID = 'c2f71bf6-7256-4428-8920-fd7b15859c2f') 
  INSERT [dbo].[CoverageItemDefinitions] ([ID], [Code], [Title], [IsTop5], [IsTop20], [SortingIndex], [Description], [Group_ID], [Logo], [ValueType], [ValueProvider], [FiltrationSteps], [Category], [IsTopItem], [UsedForPlanScore], [Profile], [HealthProfile], [IsShowInFormFill], [TopItemSortingIndex]) 
  VALUES (N'c2f71bf6-7256-4428-8920-fd7b15859c2f', N'TripCancellation', N'Cancelled trip', 1, 0, 14, N'This covers you when you cancel your trip – but only if you have a good reason! Illness, death in the family, or other situations you couldn’t possibly plan for are usually covered. ', N'8c60abf5-8480-43cd-8bf3-3a93897a3df5', N'https://gobear0sg0prod.blob.core.windows.net/coveragelogos/trip_cancellation.svg?rnd=1418960347462', 0, NULL, N'0,2000,5000,10000,15000,20000', NULL, 0, 0, N'All', 1, 0, 5)

