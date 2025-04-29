package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dataStr := `[
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "qJViJZUBV5X_QYDfAJwV",
                "_score": 1.0047282,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 837360,
                        "total_size": 238297819454,
                        "group_by_agent": [
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 1832,
                                "total_size": 145569076
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 1180,
                                "total_size": 294380805
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 1277,
                                "total_size": 642840555
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 224,
                                "total_size": 149072680
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 16450,
                                "total_size": 1396264664
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 1980,
                                "total_size": 608002175
                            },
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 1988,
                                "total_size": 500315527
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 40,
                                "total_size": 689091
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 200,
                                "total_size": 11426530
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 12,
                                "total_size": 1353292
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 44,
                                "total_size": 7873788
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 1976,
                                "total_size": 6000252194
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 165,
                                "total_size": 12229346
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 4224,
                                "total_size": 1445293361
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 74,
                                "total_size": 52759830
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 419,
                                "total_size": 178964688
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 2273,
                                "total_size": 681106880
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 1952,
                                "total_size": 4865712028
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 3,
                                "total_size": 3545
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 2398,
                                "total_size": 575395027
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 250,
                                "total_size": 21034249
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 2043,
                                "total_size": 8316593280
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 481,
                                "total_size": 299656570
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 327,
                                "total_size": 5955097
                            },
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 2451,
                                "total_size": 198436877
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 76,
                                "total_size": 6394257
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 1823,
                                "total_size": 146407452
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 1748,
                                "total_size": 425687368
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 976,
                                "total_size": 6096847570
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 144,
                                "total_size": 3662025
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 90,
                                "total_size": 24178683
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 8693,
                                "total_size": 886358819
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 2533,
                                "total_size": 636739461
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 800,
                                "total_size": 4990482865
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 1216,
                                "total_size": 86538429
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 44,
                                "total_size": 66076021
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 1080,
                                "total_size": 99032651
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 1304,
                                "total_size": 341726026
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 692,
                                "total_size": 383158874
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 15898,
                                "total_size": 1049539793
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 69,
                                "total_size": 38249483
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 7471,
                                "total_size": 673546784
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 317,
                                "total_size": 192682958
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 17542,
                                "total_size": 1725102086
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 342,
                                "total_size": 327243224
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 174,
                                "total_size": 63884880
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 988,
                                "total_size": 74528056
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 1076,
                                "total_size": 595689677
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 240,
                                "total_size": 65691253
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 872,
                                "total_size": 229270810
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 100,
                                "total_size": 11891503
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 390,
                                "total_size": 155694935
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 2520,
                                "total_size": 753196185
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 72,
                                "total_size": 96778
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 404,
                                "total_size": 314292807
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 1226,
                                "total_size": 312139268
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 1233,
                                "total_size": 315180003
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 35,
                                "total_size": 16628584
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 1552,
                                "total_size": 1573878417
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 1288,
                                "total_size": 122596605
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 2750,
                                "total_size": 1120208974
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 12705,
                                "total_size": 1181378896
                            },
                            {
                                "rc_agent": "456d9bf9-a53a-4786-9c2c-24cdd5fd6771",
                                "total_count": 2344,
                                "total_size": 1039332605
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 13776,
                                "total_size": 1311441438
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 1385,
                                "total_size": 124283746
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 4296,
                                "total_size": 912374432
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 418,
                                "total_size": 131073775
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 848,
                                "total_size": 91284869
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 1192,
                                "total_size": 282506368
                            },
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 2553,
                                "total_size": 169035537
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 16958,
                                "total_size": 1556976234
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 2462,
                                "total_size": 627352245
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 1067,
                                "total_size": 363910136
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 1390,
                                "total_size": 460576642
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 136,
                                "total_size": 1084262662
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 2897,
                                "total_size": 383959098
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 96,
                                "total_size": 18275586
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 128,
                                "total_size": 106094559
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 966,
                                "total_size": 286834375
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 3228,
                                "total_size": 1539417452
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 24443,
                                "total_size": 2170408484
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 1510,
                                "total_size": 115003656
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 1530,
                                "total_size": 108479158
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 15826,
                                "total_size": 1551229362
                            },
                            {
                                "rc_agent": "7fe8e59a-b5c1-4e36-a9f4-b526023335fc",
                                "total_count": 27270,
                                "total_size": 2891546143
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 8,
                                "total_size": 1775483
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 96,
                                "total_size": 10321346
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 3200,
                                "total_size": 6446894659
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 742,
                                "total_size": 489537171
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 297,
                                "total_size": 80706948
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 2752,
                                "total_size": 3715829197
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 312,
                                "total_size": 751875117
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 88,
                                "total_size": 67362962
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 102,
                                "total_size": 21964054
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 13810,
                                "total_size": 926072099
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 407,
                                "total_size": 316037423
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 141,
                                "total_size": 305471994
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 868,
                                "total_size": 108213013
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 72,
                                "total_size": 11674741
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 112,
                                "total_size": 782339466
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 1328,
                                "total_size": 301543675
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 38,
                                "total_size": 1986354
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 144,
                                "total_size": 1751017358
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 88,
                                "total_size": 3845164
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 494,
                                "total_size": 22140476
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 1394,
                                "total_size": 54445658
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 120,
                                "total_size": 2143643
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 1157,
                                "total_size": 53088383
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 17398,
                                "total_size": 1539981062
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 2441,
                                "total_size": 400973156
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 16499,
                                "total_size": 1558244524
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 2383,
                                "total_size": 193238608
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 138,
                                "total_size": 2419242
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 1130,
                                "total_size": 375942972
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 1228,
                                "total_size": 332583035
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 1848,
                                "total_size": 427257211
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 1084,
                                "total_size": 16491126
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 96,
                                "total_size": 33225910
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 512,
                                "total_size": 87547504
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 404,
                                "total_size": 47063022
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 265,
                                "total_size": 72068337
                            },
                            {
                                "rc_agent": "0143a74d-9e52-4478-aba0-f0853a0d9f02",
                                "total_count": 1325,
                                "total_size": 308727809
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 2410,
                                "total_size": 113704667
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 376,
                                "total_size": 294090063
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 232,
                                "total_size": 108476541
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 1773,
                                "total_size": 700114453
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 404,
                                "total_size": 47954940
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 3964,
                                "total_size": 2834171108
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 1456,
                                "total_size": 596579367
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 1574,
                                "total_size": 612145679
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 96,
                                "total_size": 4348187
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 240,
                                "total_size": 12934233
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 824,
                                "total_size": 663929549
                            },
                            {
                                "rc_agent": "e5bfa1ac-6282-4695-b99f-b948c9244adc",
                                "total_count": 9546,
                                "total_size": 850366280
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 891,
                                "total_size": 25768778
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 1336,
                                "total_size": 1034777218
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 3,
                                "total_size": 16902236
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 1528,
                                "total_size": 1462835654
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 335,
                                "total_size": 6573084
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 8,
                                "total_size": 4111399
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 673,
                                "total_size": 50193473
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 5805,
                                "total_size": 535569966
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 2200,
                                "total_size": 568575623
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 3418,
                                "total_size": 6312755714
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 1008,
                                "total_size": 305507051
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 950,
                                "total_size": 37252756
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 1680,
                                "total_size": 127285119
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 1484,
                                "total_size": 103633713
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 1318,
                                "total_size": 276814381
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 1349,
                                "total_size": 514386799
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 1242,
                                "total_size": 35721950
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 400,
                                "total_size": 54794339
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 15347,
                                "total_size": 1567003005
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 17952,
                                "total_size": 1556374462
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 18,
                                "total_size": 1706154
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 12965,
                                "total_size": 572335632
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 812,
                                "total_size": 95487690
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 232,
                                "total_size": 55467706
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 3990,
                                "total_size": 2319518404
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 1300,
                                "total_size": 3682831406
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 6057,
                                "total_size": 571686634
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 1670,
                                "total_size": 137450688
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 30,
                                "total_size": 663760898
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 2261,
                                "total_size": 1181363534
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 72,
                                "total_size": 12242750
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 2762,
                                "total_size": 1375085443
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 5464,
                                "total_size": 3858476458
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 2968,
                                "total_size": 895523884
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 5344,
                                "total_size": 3525641150
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 237,
                                "total_size": 10835236
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 14417,
                                "total_size": 1343593037
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 3921,
                                "total_size": 1847152521
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 286,
                                "total_size": 166563312
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 131,
                                "total_size": 388744091
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 40,
                                "total_size": 1468024
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 12667,
                                "total_size": 1094982724
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 812,
                                "total_size": 112513011
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 14514,
                                "total_size": 1152436119
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 9791,
                                "total_size": 796758048
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 193,
                                "total_size": 4015617
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 97,
                                "total_size": 71243984
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 3024,
                                "total_size": 1702186876
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 52,
                                "total_size": 5233793
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 4056,
                                "total_size": 2930626150
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 1203,
                                "total_size": 261653416
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 1288,
                                "total_size": 41006138
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 16195,
                                "total_size": 1467114845
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 1984,
                                "total_size": 169276689
                            },
                            {
                                "rc_agent": "0c98c7b5-f0bf-4b03-a667-a754e5acefd3",
                                "total_count": 7360,
                                "total_size": 647133023
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 82,
                                "total_size": 2084881
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 232,
                                "total_size": 24461578
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 937,
                                "total_size": 490019461
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 10720,
                                "total_size": 840018084
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 1092,
                                "total_size": 377295467
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 744,
                                "total_size": 346597474
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 1960,
                                "total_size": 5924206033
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 13,
                                "total_size": 10455324
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 37,
                                "total_size": 4728488
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 1769,
                                "total_size": 598097433
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 24,
                                "total_size": 760482981
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 1300,
                                "total_size": 171975000
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 88,
                                "total_size": 68889629
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 82,
                                "total_size": 2109745
                            },
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 4525,
                                "total_size": 457518439
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 2275,
                                "total_size": 588526889
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 410,
                                "total_size": 246327575
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 1304,
                                "total_size": 212365582
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 3872,
                                "total_size": 11493511689
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 1616,
                                "total_size": 125689565
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 1512,
                                "total_size": 155067712
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 96,
                                "total_size": 10971217
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 108,
                                "total_size": 9728126
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 3136,
                                "total_size": 1472105315
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 126,
                                "total_size": 10597014
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 36,
                                "total_size": 9799988
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 1206,
                                "total_size": 61405198
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 1640,
                                "total_size": 147774839
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 22502,
                                "total_size": 2095684148
                            },
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 2178,
                                "total_size": 217603096
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 1439,
                                "total_size": 105482626
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 1667,
                                "total_size": 124328338
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 20812,
                                "total_size": 1854392511
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 858,
                                "total_size": 856547668
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 11,
                                "total_size": 585993
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 2486,
                                "total_size": 4451784066
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 1441,
                                "total_size": 296512807
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 76,
                                "total_size": 5201942
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 1648,
                                "total_size": 1144355548
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 400,
                                "total_size": 93509943
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 16693,
                                "total_size": 1569176348
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 1527,
                                "total_size": 594917294
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 992,
                                "total_size": 217943293
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 724,
                                "total_size": 171361784
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 306,
                                "total_size": 13407591
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 392,
                                "total_size": 241149901
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 69,
                                "total_size": 393093
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 1448,
                                "total_size": 522067691
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 14,
                                "total_size": 787140
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 4733,
                                "total_size": 3236446456
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 12361,
                                "total_size": 769705789
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 2741,
                                "total_size": 1767817088
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 32640,
                                "total_size": 3108733396
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 858,
                                "total_size": 807724102
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 28592,
                                "total_size": 2776800693
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 912,
                                "total_size": 516885747
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 3555,
                                "total_size": 1987029861
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 1911,
                                "total_size": 627885909
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 1456,
                                "total_size": 92681953
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 2720,
                                "total_size": 792109900
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 17451,
                                "total_size": 1637499072
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 64,
                                "total_size": 106662937
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 1231,
                                "total_size": 4197677994
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 92,
                                "total_size": 29150775
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 831,
                                "total_size": 337219029
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 154,
                                "total_size": 1523824
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 1264,
                                "total_size": 116073714
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 1448,
                                "total_size": 109578636
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 852,
                                "total_size": 5324804267
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 1568,
                                "total_size": 132742955
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 139,
                                "total_size": 5841163
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 858,
                                "total_size": 1247398008
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 692,
                                "total_size": 25973836
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 8,
                                "total_size": 831762
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 270,
                                "total_size": 26668218
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 2029,
                                "total_size": 510768135
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 314,
                                "total_size": 24135371
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 1578,
                                "total_size": 4425433496
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 3536,
                                "total_size": 3306743508
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 80,
                                "total_size": 864223
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 5637,
                                "total_size": 3777032679
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 840,
                                "total_size": 285829717
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 240,
                                "total_size": 419854930
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 583,
                                "total_size": 112404249
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 49,
                                "total_size": 159923659
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 564,
                                "total_size": 347240370
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 1776,
                                "total_size": 420732674
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 1396,
                                "total_size": 18604475
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 48,
                                "total_size": 1906364
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 393,
                                "total_size": 15948608
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 957,
                                "total_size": 386441761
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 128,
                                "total_size": 80844449
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 1552,
                                "total_size": 1015905393
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 1336,
                                "total_size": 1063332543
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 18,
                                "total_size": 9096076
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 3947,
                                "total_size": 349892436
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 250,
                                "total_size": 8332851
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 96,
                                "total_size": 1133808
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 1763,
                                "total_size": 730779925
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 1167,
                                "total_size": 467263360
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 36,
                                "total_size": 3528118
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 618,
                                "total_size": 233457192
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 660,
                                "total_size": 479739854
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 863,
                                "total_size": 23673147
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 1168,
                                "total_size": 1036756413
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 1168,
                                "total_size": 804870869
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 1048,
                                "total_size": 701055639
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 3752,
                                "total_size": 1807788480
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 2297,
                                "total_size": 1118837520
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 1349,
                                "total_size": 98298268
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 77,
                                "total_size": 2339312
                            }
                        ]
                    },
                    "time": "2025-02-20T01:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "rZViJZUBV5X_QYDfAJwV",
                "_score": 1.0047282,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 524342,
                        "total_size": 158769304710,
                        "group_by_agent": [
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 2350,
                                "total_size": 219486732
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 713,
                                "total_size": 240554345
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 1264,
                                "total_size": 505170591
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 671,
                                "total_size": 90766275
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 536,
                                "total_size": 163001062
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 28,
                                "total_size": 480668
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 6,
                                "total_size": 587126
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 807,
                                "total_size": 61821723
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 442,
                                "total_size": 343510466
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 863,
                                "total_size": 138981213
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 199,
                                "total_size": 194512818
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 471,
                                "total_size": 120366948
                            },
                            {
                                "rc_agent": "7fe8e59a-b5c1-4e36-a9f4-b526023335fc",
                                "total_count": 11186,
                                "total_size": 1200590443
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 114,
                                "total_size": 2366194
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 2030,
                                "total_size": 1432321451
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 37,
                                "total_size": 944046
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 111,
                                "total_size": 2063216
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 1748,
                                "total_size": 156976869
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 924,
                                "total_size": 69853422
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 654,
                                "total_size": 399472512
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 257,
                                "total_size": 203234029
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 1542,
                                "total_size": 109767463
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 53,
                                "total_size": 16505688
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 77,
                                "total_size": 805858585
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 438,
                                "total_size": 17080975
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 847,
                                "total_size": 302061092
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 748,
                                "total_size": 21563978
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 702,
                                "total_size": 154494658
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 990,
                                "total_size": 494786896
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 335,
                                "total_size": 208382721
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 34,
                                "total_size": 867587
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 8,
                                "total_size": 6574969
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 18382,
                                "total_size": 1795776671
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 2258,
                                "total_size": 2114074419
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 248,
                                "total_size": 184565246
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 1197,
                                "total_size": 255536649
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 955,
                                "total_size": 740371730
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 444,
                                "total_size": 73436356
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 435,
                                "total_size": 220215729
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 551,
                                "total_size": 149473931
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 470,
                                "total_size": 312916508
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 206,
                                "total_size": 18857088
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 2968,
                                "total_size": 1491894863
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 3888,
                                "total_size": 3220509562
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 3062,
                                "total_size": 4076940804
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 1520,
                                "total_size": 113348718
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 741,
                                "total_size": 293938557
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 1316,
                                "total_size": 117245897
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 515,
                                "total_size": 42739423
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 929,
                                "total_size": 71418793
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 126,
                                "total_size": 4219687
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 304,
                                "total_size": 200914990
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 601,
                                "total_size": 579953153
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 1239,
                                "total_size": 2288235742
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 1321,
                                "total_size": 3995743630
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 186,
                                "total_size": 110290404
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 1644,
                                "total_size": 783815701
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 74,
                                "total_size": 106796488
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 3334,
                                "total_size": 997134654
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 22,
                                "total_size": 966071
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 352,
                                "total_size": 32297911
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 82,
                                "total_size": 2070843
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 4308,
                                "total_size": 2919454626
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 1445,
                                "total_size": 589036867
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 2712,
                                "total_size": 1536818637
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 597,
                                "total_size": 44212838
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 6123,
                                "total_size": 413855930
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 202,
                                "total_size": 39329527
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 194,
                                "total_size": 52736002
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 592,
                                "total_size": 98240274
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 1139,
                                "total_size": 375263426
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 262,
                                "total_size": 146231372
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 1393,
                                "total_size": 360394119
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 1329,
                                "total_size": 925978290
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 260,
                                "total_size": 12208956
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 11706,
                                "total_size": 1082034803
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 1086,
                                "total_size": 90705831
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 1327,
                                "total_size": 99176996
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 708,
                                "total_size": 710737329
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 2834,
                                "total_size": 1651356259
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 116,
                                "total_size": 8541540
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 7,
                                "total_size": 359359
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 3,
                                "total_size": 3545
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 7470,
                                "total_size": 678463592
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 9668,
                                "total_size": 916209712
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 16874,
                                "total_size": 1515321429
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 937,
                                "total_size": 828149772
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 360,
                                "total_size": 7976236
                            },
                            {
                                "rc_agent": "e5bfa1ac-6282-4695-b99f-b948c9244adc",
                                "total_count": 848,
                                "total_size": 75975005
                            },
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 1270,
                                "total_size": 323915857
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 80,
                                "total_size": 5638225
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 257,
                                "total_size": 96254660
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 314,
                                "total_size": 177528179
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 1352,
                                "total_size": 4132877762
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 1580,
                                "total_size": 640617056
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 54,
                                "total_size": 17286165
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 434,
                                "total_size": 35588426
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 719,
                                "total_size": 54335584
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 703,
                                "total_size": 1025931698
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 925,
                                "total_size": 618065202
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 31,
                                "total_size": 3708410
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 1358,
                                "total_size": 405249204
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 221,
                                "total_size": 131616562
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 652,
                                "total_size": 18039150
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 288,
                                "total_size": 28873841
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 376,
                                "total_size": 230970559
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 441,
                                "total_size": 271609351
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 46,
                                "total_size": 74143213
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 1865,
                                "total_size": 1199155336
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 752,
                                "total_size": 204898104
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 370,
                                "total_size": 173501974
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 599,
                                "total_size": 61102923
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 853,
                                "total_size": 5507179653
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 3038,
                                "total_size": 643318138
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 50,
                                "total_size": 37779182
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 1041,
                                "total_size": 830160364
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 6,
                                "total_size": 3051297
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 2134,
                                "total_size": 1204148460
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 3060,
                                "total_size": 291807931
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 715,
                                "total_size": 182157994
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 190,
                                "total_size": 3710833
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 746,
                                "total_size": 177004612
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 233,
                                "total_size": 170052632
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 604,
                                "total_size": 607922215
                            },
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 2669,
                                "total_size": 215299373
                            },
                            {
                                "rc_agent": "456d9bf9-a53a-4786-9c2c-24cdd5fd6771",
                                "total_count": 1150,
                                "total_size": 513759654
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 409,
                                "total_size": 13100844
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 1189,
                                "total_size": 296688796
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 241,
                                "total_size": 142826827
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 10853,
                                "total_size": 884488916
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 723,
                                "total_size": 77528130
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 10833,
                                "total_size": 1021601056
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 556,
                                "total_size": 132222414
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 560,
                                "total_size": 45368769
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 13146,
                                "total_size": 1209677813
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 102,
                                "total_size": 12134477
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 612,
                                "total_size": 178809860
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 3236,
                                "total_size": 280415582
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 49,
                                "total_size": 842290
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 726,
                                "total_size": 98488394
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 769,
                                "total_size": 3081942551
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 37,
                                "total_size": 178313348
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 988,
                                "total_size": 376799353
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 21,
                                "total_size": 1106335
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 650,
                                "total_size": 25366446
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 2433,
                                "total_size": 1735384644
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 347,
                                "total_size": 23537635
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 6921,
                                "total_size": 628877609
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 11348,
                                "total_size": 1065816810
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 312,
                                "total_size": 916081631
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 1114,
                                "total_size": 445138678
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 15,
                                "total_size": 838263
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 1989,
                                "total_size": 504086528
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 22,
                                "total_size": 20966449
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 831,
                                "total_size": 88196708
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 12,
                                "total_size": 402747195
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 467,
                                "total_size": 241672050
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 2282,
                                "total_size": 203919979
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 13032,
                                "total_size": 1199404546
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 8632,
                                "total_size": 876553618
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 7972,
                                "total_size": 498347440
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 1216,
                                "total_size": 473111171
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 16,
                                "total_size": 120773
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 203,
                                "total_size": 55216246
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 1471,
                                "total_size": 688812815
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 1397,
                                "total_size": 425877409
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 1692,
                                "total_size": 446400591
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 126,
                                "total_size": 46137811
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 63,
                                "total_size": 639379
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 319,
                                "total_size": 2011998661
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 28,
                                "total_size": 9639720
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 9686,
                                "total_size": 993391171
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 2147,
                                "total_size": 212965494
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 51,
                                "total_size": 2235659
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 105,
                                "total_size": 1236627
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 3567,
                                "total_size": 7229956946
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 214,
                                "total_size": 171184606
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 9096,
                                "total_size": 1008843640
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 67,
                                "total_size": 3597413
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 1950,
                                "total_size": 199625794
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 424,
                                "total_size": 43432900
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 150,
                                "total_size": 2684374
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 770,
                                "total_size": 96918331
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 46,
                                "total_size": 500699
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 231,
                                "total_size": 10312528
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 463,
                                "total_size": 106764313
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 258,
                                "total_size": 18852977
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 638,
                                "total_size": 114343807
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 305,
                                "total_size": 148496652
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 1316,
                                "total_size": 3278092946
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 820,
                                "total_size": 225915278
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 2434,
                                "total_size": 203070007
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 1394,
                                "total_size": 358672510
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 2296,
                                "total_size": 194727765
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 137,
                                "total_size": 13455238
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 228,
                                "total_size": 38187152
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 215,
                                "total_size": 16151196
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 5828,
                                "total_size": 457429944
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 387,
                                "total_size": 81952642
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 23,
                                "total_size": 674334198
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 138,
                                "total_size": 254976349
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 30,
                                "total_size": 904333
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 942,
                                "total_size": 249010801
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 1389,
                                "total_size": 124772287
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 26,
                                "total_size": 39026269
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 2432,
                                "total_size": 1132195147
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 6758,
                                "total_size": 573441462
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 12565,
                                "total_size": 994234607
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 3413,
                                "total_size": 2470563302
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 258,
                                "total_size": 11395805
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 1978,
                                "total_size": 483695862
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 15857,
                                "total_size": 1394805350
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 1315,
                                "total_size": 2299415911
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 951,
                                "total_size": 662074051
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 4926,
                                "total_size": 452066577
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 982,
                                "total_size": 68455362
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 57,
                                "total_size": 671760598
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 9833,
                                "total_size": 648602840
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 797,
                                "total_size": 42555218
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 108,
                                "total_size": 12052761
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 115,
                                "total_size": 5291293
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 4062,
                                "total_size": 360234066
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 791,
                                "total_size": 136732694
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 19779,
                                "total_size": 1912946030
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 752,
                                "total_size": 248359811
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 774,
                                "total_size": 227258885
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 663,
                                "total_size": 164976968
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 210,
                                "total_size": 8424643
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 406,
                                "total_size": 128670592
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 197,
                                "total_size": 80967950
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 8571,
                                "total_size": 738696563
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 18529,
                                "total_size": 1764022104
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 958,
                                "total_size": 48386126
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 17,
                                "total_size": 2164398
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 831,
                                "total_size": 289781529
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 54,
                                "total_size": 42559466
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 51,
                                "total_size": 2313980
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 9,
                                "total_size": 1023402
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 925,
                                "total_size": 67857393
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 751,
                                "total_size": 160936558
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 3136,
                                "total_size": 9212363150
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 507,
                                "total_size": 1728856818
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 1353,
                                "total_size": 85301912
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 2225,
                                "total_size": 760582260
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 645,
                                "total_size": 122937389
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 712,
                                "total_size": 293253543
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 144,
                                "total_size": 109510579
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 70,
                                "total_size": 125947135
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 1525,
                                "total_size": 105412384
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 488,
                                "total_size": 191847896
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 1386,
                                "total_size": 361196895
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 419,
                                "total_size": 136659405
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 716,
                                "total_size": 732422343
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 1010,
                                "total_size": 2761262172
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 336,
                                "total_size": 76777561
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 1668,
                                "total_size": 791425617
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 614,
                                "total_size": 156615065
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 582,
                                "total_size": 209713212
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 935,
                                "total_size": 212623512
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 38,
                                "total_size": 216602
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 635,
                                "total_size": 174135857
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 10047,
                                "total_size": 943525259
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 4083,
                                "total_size": 361924190
                            },
                            {
                                "rc_agent": "0143a74d-9e52-4478-aba0-f0853a0d9f02",
                                "total_count": 592,
                                "total_size": 140147449
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 3323,
                                "total_size": 1653996719
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 672,
                                "total_size": 74852491
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 623,
                                "total_size": 263461061
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 176,
                                "total_size": 443258040
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 498,
                                "total_size": 52723438
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 971,
                                "total_size": 796899056
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 878,
                                "total_size": 208905136
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 2047,
                                "total_size": 270386203
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 15,
                                "total_size": 594516
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 849,
                                "total_size": 254555974
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 2122,
                                "total_size": 217740587
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 417,
                                "total_size": 179804692
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 812,
                                "total_size": 11480752
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 2561,
                                "total_size": 1687078853
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 404,
                                "total_size": 224919099
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 512,
                                "total_size": 50612870
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 782,
                                "total_size": 4945124406
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 2,
                                "total_size": 8920951
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 2014,
                                "total_size": 1423548041
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 96,
                                "total_size": 1323331087
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 168,
                                "total_size": 100876088
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 300,
                                "total_size": 67591061
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 984,
                                "total_size": 128022981
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 59,
                                "total_size": 6757733
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 599,
                                "total_size": 392795489
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 863,
                                "total_size": 206355368
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 422,
                                "total_size": 218901343
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 760,
                                "total_size": 21918354
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 1085,
                                "total_size": 51407920
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 422,
                                "total_size": 142785756
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 403,
                                "total_size": 15112700
                            }
                        ]
                    },
                    "time": "2025-02-20T04:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "d5XJJJUBLBoVSWi0Gx7k",
                "_score": 1.0047282,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T18:00:00.000000Z",
                    "indices": {
                        "total_count": 229,
                        "total_size": 15232330,
                        "group_by_agent": [
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 229,
                                "total_size": 15232330
                            }
                        ]
                    },
                    "time": "2025-02-20T09:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "YkYAJZUBT5H4H6HfCzau",
                "_score": 1.0047282,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 228,
                        "total_size": 21718172,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 228,
                                "total_size": 21718172
                            }
                        ]
                    },
                    "time": "2025-02-20T06:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "2rySJJUBk0bmQQGuSRjp",
                "_score": 1.0047282,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T17:00:00.000000Z",
                    "indices": {
                        "total_count": 177,
                        "total_size": 14487592,
                        "group_by_agent": [
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 177,
                                "total_size": 14487592
                            }
                        ]
                    },
                    "time": "2025-02-20T05:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "XtxJJpUBWlvdGmPA0qZ8",
                "_score": 1.0047282,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-21T01:00:00.000000Z",
                    "indices": {
                        "total_count": 3,
                        "total_size": 269561,
                        "group_by_agent": [
                            {
                                "rc_agent": "c18624e7-e7f5-48c1-9352-4965a0f8c59e",
                                "total_count": 3,
                                "total_size": 269561
                            }
                        ]
                    },
                    "time": "2025-02-20T06:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "eJXJJJUBLBoVSWi0Gx7k",
                "_score": 1.0046839,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T18:00:00.000000Z",
                    "indices": {
                        "total_count": 396,
                        "total_size": 26565369,
                        "group_by_agent": [
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 396,
                                "total_size": 26565369
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "Y0YAJZUBT5H4H6HfCzau",
                "_score": 1.0046839,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 287,
                        "total_size": 28635213,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 287,
                                "total_size": 28635213
                            }
                        ]
                    },
                    "time": "2025-02-20T05:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "bZVbJJUBLBoVSWi0Sx7K",
                "_score": 1.0046839,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T16:00:00.000000Z",
                    "indices": {
                        "total_count": 2,
                        "total_size": 545954,
                        "group_by_agent": [
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 2,
                                "total_size": 545954
                            }
                        ]
                    },
                    "time": "2025-02-20T08:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "P9xtJZUBWlvdGmPA86Y9",
                "_score": 1.0046839,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T21:00:00.000000Z",
                    "indices": {
                        "total_count": 45,
                        "total_size": 21787072,
                        "group_by_agent": [
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 12,
                                "total_size": 3266345
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 33,
                                "total_size": 18520727
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "zrySJJUBk0bmQQGuSRjp",
                "_score": 1.0046839,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T17:00:00.000000Z",
                    "indices": {
                        "total_count": 55,
                        "total_size": 4605346,
                        "group_by_agent": [
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 55,
                                "total_size": 4605346
                            }
                        ]
                    },
                    "time": "2025-02-20T08:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "07ySJJUBk0bmQQGuSRjp",
                "_score": 1.0046839,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T17:00:00.000000Z",
                    "indices": {
                        "total_count": 71,
                        "total_size": 5774861,
                        "group_by_agent": [
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 71,
                                "total_size": 5774861
                            }
                        ]
                    },
                    "time": "2025-02-20T04:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "XdxJJpUBWlvdGmPA0qZ8",
                "_score": 1.0046839,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-21T01:00:00.000000Z",
                    "indices": {
                        "total_count": 2,
                        "total_size": 174052,
                        "group_by_agent": [
                            {
                                "rc_agent": "c18624e7-e7f5-48c1-9352-4965a0f8c59e",
                                "total_count": 2,
                                "total_size": 174052
                            }
                        ]
                    },
                    "time": "2025-02-20T08:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "mZViJZUBV5X_QYDfAJwV",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 125271,
                        "total_size": 51044480413,
                        "group_by_agent": [
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 141,
                                "total_size": 13833016
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 114,
                                "total_size": 9163297
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 196,
                                "total_size": 49690679
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 298,
                                "total_size": 195537124
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 12,
                                "total_size": 200982
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 348,
                                "total_size": 27645151
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 242,
                                "total_size": 70378869
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 910,
                                "total_size": 310874299
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 356,
                                "total_size": 140894936
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 984,
                                "total_size": 126354509
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 349,
                                "total_size": 11081275
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 204,
                                "total_size": 51428808
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 150,
                                "total_size": 1989019
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 149,
                                "total_size": 119262488
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 99,
                                "total_size": 61547494
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 23,
                                "total_size": 17914050
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 21,
                                "total_size": 531021
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 972,
                                "total_size": 99791608
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 152,
                                "total_size": 60394060
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 293,
                                "total_size": 48590299
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 107,
                                "total_size": 51096614
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 146,
                                "total_size": 3573798
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 439,
                                "total_size": 290986071
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 323,
                                "total_size": 858789725
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 183,
                                "total_size": 29148822
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 347,
                                "total_size": 39936039
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 206,
                                "total_size": 116724456
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 3,
                                "total_size": 319735
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 854,
                                "total_size": 389823362
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 132,
                                "total_size": 30344511
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 254,
                                "total_size": 11738550
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 204,
                                "total_size": 14214526
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 304,
                                "total_size": 1266780654
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 366,
                                "total_size": 361839652
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 278,
                                "total_size": 104472206
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 374,
                                "total_size": 126856377
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 408,
                                "total_size": 96837943
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 51,
                                "total_size": 108398248
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 305,
                                "total_size": 270318588
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 181,
                                "total_size": 20624369
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 403,
                                "total_size": 36530358
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 179,
                                "total_size": 12827039
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 569,
                                "total_size": 94713781
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 15,
                                "total_size": 1587405
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 448,
                                "total_size": 39432970
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 383,
                                "total_size": 235052253
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 1,
                                "total_size": 129403
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 543,
                                "total_size": 139737636
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 273,
                                "total_size": 63647329
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 170,
                                "total_size": 56282379
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 101,
                                "total_size": 80011305
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 2437,
                                "total_size": 230114162
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 458,
                                "total_size": 105875068
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 81,
                                "total_size": 2986376
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 1927,
                                "total_size": 163715065
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 346,
                                "total_size": 29187430
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 157,
                                "total_size": 51016367
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 630,
                                "total_size": 255371703
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 18,
                                "total_size": 539014
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 915,
                                "total_size": 709065324
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 587,
                                "total_size": 28546779
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 15,
                                "total_size": 144027531
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 159,
                                "total_size": 12410187
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 2863,
                                "total_size": 256046474
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 451,
                                "total_size": 44388565
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 469,
                                "total_size": 263187783
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 35,
                                "total_size": 19471213
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 313,
                                "total_size": 31535506
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 241,
                                "total_size": 821803734
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 47,
                                "total_size": 19317965
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 22,
                                "total_size": 7310464
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 4,
                                "total_size": 3239733
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 581,
                                "total_size": 105288600
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 654,
                                "total_size": 438996945
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 203,
                                "total_size": 51540459
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 464,
                                "total_size": 44103398
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 1984,
                                "total_size": 124521392
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 279,
                                "total_size": 197311945
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 302,
                                "total_size": 16340269
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 337,
                                "total_size": 42854100
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 321,
                                "total_size": 12577922
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 146,
                                "total_size": 91457760
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 437,
                                "total_size": 120655869
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 26,
                                "total_size": 1226812
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 377,
                                "total_size": 127585202
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 451,
                                "total_size": 133237427
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 543,
                                "total_size": 166769372
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 511,
                                "total_size": 259595596
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 2922,
                                "total_size": 283601778
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 875,
                                "total_size": 531024188
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 343,
                                "total_size": 118792259
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 98,
                                "total_size": 1849356
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 918,
                                "total_size": 84011381
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 353,
                                "total_size": 22813303
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 142,
                                "total_size": 4053137
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 44,
                                "total_size": 4297218
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 48,
                                "total_size": 29227465
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 439,
                                "total_size": 255365425
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 127,
                                "total_size": 40590675
                            },
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 2698,
                                "total_size": 267612381
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 249,
                                "total_size": 55347473
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 47,
                                "total_size": 828799
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 381,
                                "total_size": 978204270
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 2,
                                "total_size": 16178236
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 518,
                                "total_size": 130595097
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 390,
                                "total_size": 152428068
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 387,
                                "total_size": 284642124
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 14,
                                "total_size": 8561830
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 372,
                                "total_size": 29612415
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 197,
                                "total_size": 18032451
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 141,
                                "total_size": 13210655
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 32,
                                "total_size": 2350845
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 30,
                                "total_size": 622928
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 4,
                                "total_size": 101034
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 23,
                                "total_size": 13916451
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 328,
                                "total_size": 24032584
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 32,
                                "total_size": 591363
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 126,
                                "total_size": 46916176
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 266,
                                "total_size": 99397836
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 431,
                                "total_size": 29892603
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 10,
                                "total_size": 1135613
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 398,
                                "total_size": 1147069461
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 547,
                                "total_size": 266272388
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 3540,
                                "total_size": 343748139
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 1814,
                                "total_size": 164589410
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 9,
                                "total_size": 13502904
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 33,
                                "total_size": 1017438
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 17,
                                "total_size": 13790391
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 640,
                                "total_size": 311412012
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 341,
                                "total_size": 27637857
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 425,
                                "total_size": 126219722
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 66,
                                "total_size": 7554108
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 2032,
                                "total_size": 159228268
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 292,
                                "total_size": 101749924
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 1,
                                "total_size": 852618
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 11,
                                "total_size": 9593004
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 319,
                                "total_size": 126929265
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 364,
                                "total_size": 253131131
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 86,
                                "total_size": 57354765
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 281,
                                "total_size": 10789238
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 287,
                                "total_size": 65279344
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 227,
                                "total_size": 87327123
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 318,
                                "total_size": 84816735
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 341,
                                "total_size": 28898617
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 1208,
                                "total_size": 614481648
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 107,
                                "total_size": 7904991
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 140,
                                "total_size": 31382612
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 161,
                                "total_size": 39606012
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 1262,
                                "total_size": 1701802887
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 157,
                                "total_size": 67201897
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 468,
                                "total_size": 329950661
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 463,
                                "total_size": 31289066
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 333,
                                "total_size": 28981793
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 225,
                                "total_size": 114701263
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 302,
                                "total_size": 78206944
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 39,
                                "total_size": 103710282
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 893,
                                "total_size": 426956004
                            },
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 452,
                                "total_size": 110323552
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 688,
                                "total_size": 665925749
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 418,
                                "total_size": 37516947
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 82,
                                "total_size": 9904730
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 601,
                                "total_size": 39204103
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 475,
                                "total_size": 48190697
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 154,
                                "total_size": 15502751
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 212,
                                "total_size": 86738099
                            },
                            {
                                "rc_agent": "0143a74d-9e52-4478-aba0-f0853a0d9f02",
                                "total_count": 287,
                                "total_size": 71727572
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 389,
                                "total_size": 29100916
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 234,
                                "total_size": 60203198
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 327,
                                "total_size": 299837674
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 1737,
                                "total_size": 154816629
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 569,
                                "total_size": 397008072
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 13,
                                "total_size": 4102813
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 15,
                                "total_size": 548087
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 2,
                                "total_size": 2351
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 180,
                                "total_size": 24655409
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 144,
                                "total_size": 80287449
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 2487,
                                "total_size": 223044603
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 415,
                                "total_size": 41381405
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 318,
                                "total_size": 243453810
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 13,
                                "total_size": 1543824
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 477,
                                "total_size": 205252837
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 114,
                                "total_size": 81946797
                            },
                            {
                                "rc_agent": "456d9bf9-a53a-4786-9c2c-24cdd5fd6771",
                                "total_count": 643,
                                "total_size": 280809779
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 443,
                                "total_size": 210711786
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 1942,
                                "total_size": 85875007
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 321,
                                "total_size": 28550017
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 202,
                                "total_size": 621578003
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 413,
                                "total_size": 31066625
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 22,
                                "total_size": 127394
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 222,
                                "total_size": 37784789
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 339,
                                "total_size": 28360538
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 449,
                                "total_size": 100871544
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 51,
                                "total_size": 1014911
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 6,
                                "total_size": 24355394
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 315,
                                "total_size": 43314073
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 143,
                                "total_size": 11214094
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 441,
                                "total_size": 46466345
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 1107,
                                "total_size": 2332962097
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 23,
                                "total_size": 296676462
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 91,
                                "total_size": 59899893
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 237,
                                "total_size": 1558498909
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 319,
                                "total_size": 75574401
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 306,
                                "total_size": 37503994
                            },
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 123,
                                "total_size": 7797419
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 418,
                                "total_size": 35950105
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 68,
                                "total_size": 389399559
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 878,
                                "total_size": 75628559
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 3,
                                "total_size": 155555
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 1033,
                                "total_size": 116192105
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 381,
                                "total_size": 1218514337
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 10,
                                "total_size": 113693
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 2365,
                                "total_size": 203154530
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 419,
                                "total_size": 99362687
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 382,
                                "total_size": 148151012
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 1958,
                                "total_size": 183777795
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 90,
                                "total_size": 46419676
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 401,
                                "total_size": 320667312
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 300,
                                "total_size": 63562323
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 323,
                                "total_size": 15974806
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 191,
                                "total_size": 17443758
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 6,
                                "total_size": 193165087
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 214,
                                "total_size": 17682374
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 116,
                                "total_size": 9167594
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 981,
                                "total_size": 92270469
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 666,
                                "total_size": 59275092
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 681,
                                "total_size": 114839651
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 364,
                                "total_size": 385191886
                            },
                            {
                                "rc_agent": "cde365e1-868f-493d-9adf-7ce505eec802",
                                "total_count": 1,
                                "total_size": 55368
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 7,
                                "total_size": 518198
                            },
                            {
                                "rc_agent": "e5bfa1ac-6282-4695-b99f-b948c9244adc",
                                "total_count": 86,
                                "total_size": 7874125
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 478,
                                "total_size": 115482392
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 34,
                                "total_size": 587431820
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 584,
                                "total_size": 435460490
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 10,
                                "total_size": 72816
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 288,
                                "total_size": 585636450
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 117,
                                "total_size": 208092300
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 6,
                                "total_size": 150717
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 1239,
                                "total_size": 329516610
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 412,
                                "total_size": 314857983
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 511,
                                "total_size": 352035831
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 2364,
                                "total_size": 719159671
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 93,
                                "total_size": 4045311
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 7,
                                "total_size": 774286
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 1218,
                                "total_size": 113421813
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 94,
                                "total_size": 40556362
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 244,
                                "total_size": 73893983
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 328,
                                "total_size": 24305811
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 321,
                                "total_size": 70732931
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 301,
                                "total_size": 148402244
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 2184,
                                "total_size": 196574128
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 403,
                                "total_size": 121944046
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 269,
                                "total_size": 73426875
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 180,
                                "total_size": 68444901
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 290,
                                "total_size": 184829485
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 546,
                                "total_size": 948628052
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 30,
                                "total_size": 358773
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 894,
                                "total_size": 2767288860
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 283,
                                "total_size": 1747795638
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 40,
                                "total_size": 12547306
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 23,
                                "total_size": 36037916
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 66,
                                "total_size": 2659841
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 8,
                                "total_size": 314587
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 104,
                                "total_size": 20401605
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 279,
                                "total_size": 12389338
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 497,
                                "total_size": 103651019
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 203,
                                "total_size": 5528202
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 1254,
                                "total_size": 853479059
                            },
                            {
                                "rc_agent": "bace3ddf-e238-4660-9232-4d661fe426bf",
                                "total_count": 9,
                                "total_size": 735868
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 356,
                                "total_size": 27869007
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 6,
                                "total_size": 66415860
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 244,
                                "total_size": 7029456
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 338,
                                "total_size": 492690569
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 98,
                                "total_size": 5860510
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 7,
                                "total_size": 400764
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 175,
                                "total_size": 39710455
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 24,
                                "total_size": 246088
                            },
                            {
                                "rc_agent": "7fe8e59a-b5c1-4e36-a9f4-b526023335fc",
                                "total_count": 542,
                                "total_size": 63659671
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 23,
                                "total_size": 983981
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 322,
                                "total_size": 314212080
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 543,
                                "total_size": 436675232
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 16,
                                "total_size": 24546893
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 293,
                                "total_size": 161597616
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 140,
                                "total_size": 14361381
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 499,
                                "total_size": 123264898
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 25,
                                "total_size": 1153459
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 2159,
                                "total_size": 1799584166
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 440,
                                "total_size": 84125653
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 298,
                                "total_size": 232304843
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 50,
                                "total_size": 2286423
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 337,
                                "total_size": 93517950
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 256,
                                "total_size": 77707362
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 3089,
                                "total_size": 284534607
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 490,
                                "total_size": 35931572
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 3,
                                "total_size": 167224
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 312,
                                "total_size": 102244266
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 73,
                                "total_size": 67543832
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 516,
                                "total_size": 146249772
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 595,
                                "total_size": 350922361
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "rpViJZUBV5X_QYDfAJwV",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 389701,
                        "total_size": 107514018704,
                        "group_by_agent": [
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 1250,
                                "total_size": 626542465
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 10532,
                                "total_size": 1034663333
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 435,
                                "total_size": 32815753
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 877,
                                "total_size": 818146872
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 25,
                                "total_size": 2677825
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 863,
                                "total_size": 502520660
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 6099,
                                "total_size": 625008862
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 1471,
                                "total_size": 435051909
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 449,
                                "total_size": 104126240
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 475,
                                "total_size": 178746120
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 88,
                                "total_size": 28500850
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 478,
                                "total_size": 13091300
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 194,
                                "total_size": 283262213
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 620,
                                "total_size": 17811904
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 254,
                                "total_size": 183289833
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 477,
                                "total_size": 331289330
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 256,
                                "total_size": 28737351
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 8825,
                                "total_size": 795595251
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 644,
                                "total_size": 105407350
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 84,
                                "total_size": 26243393
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 78,
                                "total_size": 93111300
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 50,
                                "total_size": 38479232
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 84,
                                "total_size": 35883495
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 6540,
                                "total_size": 598060760
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 702,
                                "total_size": 236581544
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 1917,
                                "total_size": 1361243080
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 664,
                                "total_size": 29153582
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 607,
                                "total_size": 200884778
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 464,
                                "total_size": 42391026
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 1466,
                                "total_size": 925352040
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 1716,
                                "total_size": 1150627377
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 837,
                                "total_size": 53426934
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 1100,
                                "total_size": 114948458
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 259,
                                "total_size": 1621654828
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 458,
                                "total_size": 215856066
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 255,
                                "total_size": 29922540
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 248,
                                "total_size": 138354824
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 21,
                                "total_size": 31701554
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 153,
                                "total_size": 40655656
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 728,
                                "total_size": 39212483
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 8021,
                                "total_size": 742316209
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 709,
                                "total_size": 55203878
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 49,
                                "total_size": 577916
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 26,
                                "total_size": 449919
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 1003,
                                "total_size": 259818855
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 186,
                                "total_size": 143576554
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 5711,
                                "total_size": 491709857
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 7421,
                                "total_size": 716621315
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 56,
                                "total_size": 2529807
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 698,
                                "total_size": 290749893
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 776,
                                "total_size": 103971367
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 477,
                                "total_size": 35093246
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 227,
                                "total_size": 16746146
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 6640,
                                "total_size": 671503691
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 913,
                                "total_size": 65835710
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 10116,
                                "total_size": 918403654
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 62,
                                "total_size": 37805446
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 502,
                                "total_size": 199216818
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 55,
                                "total_size": 951503
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 957,
                                "total_size": 1860713551
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 201,
                                "total_size": 184817603
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 139,
                                "total_size": 133025811
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 718,
                                "total_size": 170891499
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 1404,
                                "total_size": 143674207
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 738,
                                "total_size": 56235022
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 1146,
                                "total_size": 3343637625
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 170,
                                "total_size": 6846934
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 963,
                                "total_size": 341977687
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 220,
                                "total_size": 42221773
                            },
                            {
                                "rc_agent": "c18624e7-e7f5-48c1-9352-4965a0f8c59e",
                                "total_count": 997,
                                "total_size": 87211760
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 653,
                                "total_size": 2641307656
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 9211,
                                "total_size": 866570746
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 659,
                                "total_size": 173316636
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 3,
                                "total_size": 2154969
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 635,
                                "total_size": 221707020
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 540,
                                "total_size": 184413566
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 631,
                                "total_size": 143529582
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 471,
                                "total_size": 108636057
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 440,
                                "total_size": 71063201
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 10484,
                                "total_size": 1018343734
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 732,
                                "total_size": 2042515614
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 391,
                                "total_size": 153484098
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 161,
                                "total_size": 90639136
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 1016,
                                "total_size": 108138414
                            },
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 53,
                                "total_size": 13147368
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 381,
                                "total_size": 29474982
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 1325,
                                "total_size": 1020972740
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 77,
                                "total_size": 631141324
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 969,
                                "total_size": 375631434
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 66,
                                "total_size": 618944
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 1636,
                                "total_size": 904978160
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 6787,
                                "total_size": 423513355
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 696,
                                "total_size": 9252620
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 12,
                                "total_size": 679455
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 8729,
                                "total_size": 793593787
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 240,
                                "total_size": 27199955
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 305,
                                "total_size": 123549115
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 647,
                                "total_size": 194366390
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 8650,
                                "total_size": 817879776
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 962,
                                "total_size": 849516596
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 1009,
                                "total_size": 254005719
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 5537,
                                "total_size": 536266681
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 80,
                                "total_size": 23850189
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 162,
                                "total_size": 3007114
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 357,
                                "total_size": 49072794
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 5441,
                                "total_size": 552525358
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 1007,
                                "total_size": 661306629
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 4,
                                "total_size": 25012491
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 41,
                                "total_size": 441353254
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 5419,
                                "total_size": 481247074
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 2150,
                                "total_size": 458580949
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 261,
                                "total_size": 28429087
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 24,
                                "total_size": 47188739
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 441,
                                "total_size": 31771995
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 1392,
                                "total_size": 648855160
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 479,
                                "total_size": 385195383
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 288,
                                "total_size": 157754501
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 1514,
                                "total_size": 155629233
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 22,
                                "total_size": 856924
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 1043,
                                "total_size": 416835972
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 27,
                                "total_size": 1065691
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 12,
                                "total_size": 1365757
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 1223,
                                "total_size": 1140929570
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 585,
                                "total_size": 296327155
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 449,
                                "total_size": 110041630
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 1089,
                                "total_size": 277576166
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 133,
                                "total_size": 36202288
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 3944,
                                "total_size": 265976717
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 1074,
                                "total_size": 2014913225
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 135,
                                "total_size": 2630949
                            },
                            {
                                "rc_agent": "cde365e1-868f-493d-9adf-7ce505eec802",
                                "total_count": 1,
                                "total_size": 861516
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 1005,
                                "total_size": 83888936
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 64,
                                "total_size": 396954091
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 1034,
                                "total_size": 77564879
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 910,
                                "total_size": 95259953
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 460,
                                "total_size": 34217897
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 938,
                                "total_size": 1786929898
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 47,
                                "total_size": 27469084
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 2647,
                                "total_size": 3658411952
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 1129,
                                "total_size": 269777385
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 28,
                                "total_size": 303933
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 1747,
                                "total_size": 1183495746
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 20,
                                "total_size": 7839970
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 43,
                                "total_size": 1898816
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 582,
                                "total_size": 1792358879
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 350,
                                "total_size": 218438138
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 49,
                                "total_size": 15331728
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 4951,
                                "total_size": 408540210
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 952,
                                "total_size": 231278221
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 555,
                                "total_size": 133448660
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 979,
                                "total_size": 100422302
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 196,
                                "total_size": 51670358
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 960,
                                "total_size": 223780635
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 367,
                                "total_size": 33662898
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 193,
                                "total_size": 349949026
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 151,
                                "total_size": 6788794
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 39,
                                "total_size": 227984
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 1407,
                                "total_size": 229025564
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 2801,
                                "total_size": 1461259936
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 354,
                                "total_size": 26209657
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 254,
                                "total_size": 84188005
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 5102,
                                "total_size": 4260017723
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 255,
                                "total_size": 11138853
                            },
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 7108,
                                "total_size": 708405505
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 481,
                                "total_size": 1338945766
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 789,
                                "total_size": 187443608
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 4180,
                                "total_size": 423653299
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 2313,
                                "total_size": 180796876
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 969,
                                "total_size": 44784744
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 706,
                                "total_size": 361406722
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 121,
                                "total_size": 6895297
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 10608,
                                "total_size": 975540170
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 57,
                                "total_size": 4169362
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 972,
                                "total_size": 693738723
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 328,
                                "total_size": 131971669
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 124,
                                "total_size": 2176763
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 1077,
                                "total_size": 762472895
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 805,
                                "total_size": 330642432
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 929,
                                "total_size": 617266143
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 2803,
                                "total_size": 184776156
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 10,
                                "total_size": 216743146
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 1050,
                                "total_size": 291199301
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 380,
                                "total_size": 17070569
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 9782,
                                "total_size": 773035306
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 24,
                                "total_size": 20700119
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 1094,
                                "total_size": 283150867
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 392,
                                "total_size": 66919359
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 907,
                                "total_size": 607144082
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 12,
                                "total_size": 95051
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 743,
                                "total_size": 56474624
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 5162,
                                "total_size": 1553524809
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 390,
                                "total_size": 35154097
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 1003,
                                "total_size": 848458025
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 213,
                                "total_size": 209089526
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 7,
                                "total_size": 373795
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 38,
                                "total_size": 61667308
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 2174,
                                "total_size": 1371425556
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 417,
                                "total_size": 137694987
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 609,
                                "total_size": 245303544
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 1593,
                                "total_size": 539894009
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 965,
                                "total_size": 219785561
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 882,
                                "total_size": 894002900
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 20,
                                "total_size": 15872245
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 1625,
                                "total_size": 161381728
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 445,
                                "total_size": 36982222
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 1257,
                                "total_size": 374844418
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 24,
                                "total_size": 1233790
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 4,
                                "total_size": 269980
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 5076,
                                "total_size": 431089690
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 1323,
                                "total_size": 235923097
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 2785,
                                "total_size": 743431568
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 86,
                                "total_size": 31917229
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 40,
                                "total_size": 23223017
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 1532,
                                "total_size": 259466974
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 47,
                                "total_size": 2196326
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 690,
                                "total_size": 206196394
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 497,
                                "total_size": 252523069
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 51,
                                "total_size": 1066296
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 2151,
                                "total_size": 1552221126
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 463,
                                "total_size": 305597586
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 520,
                                "total_size": 141985926
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 18,
                                "total_size": 2305655
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 762,
                                "total_size": 212406571
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 233,
                                "total_size": 133225281
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 465,
                                "total_size": 152387411
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 441,
                                "total_size": 35047859
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 78,
                                "total_size": 6718849
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 29,
                                "total_size": 5717028
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 298,
                                "total_size": 13740144
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 229,
                                "total_size": 31054744
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 659,
                                "total_size": 20896663
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 55,
                                "total_size": 3726730
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 28,
                                "total_size": 723667
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 647,
                                "total_size": 18650288
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 40,
                                "total_size": 4707096
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 556,
                                "total_size": 21654083
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 655,
                                "total_size": 505423876
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 677,
                                "total_size": 89387815
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 34,
                                "total_size": 856707
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 938,
                                "total_size": 383940469
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 371,
                                "total_size": 193614529
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 1058,
                                "total_size": 782477005
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 294,
                                "total_size": 180582549
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 909,
                                "total_size": 270358754
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 518,
                                "total_size": 1766366532
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 7867,
                                "total_size": 697266526
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 1228,
                                "total_size": 715086513
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 302,
                                "total_size": 11434109
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 629,
                                "total_size": 509980541
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 682,
                                "total_size": 26737664
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 784,
                                "total_size": 71599108
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 91,
                                "total_size": 232485289
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 624,
                                "total_size": 1564138267
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 36,
                                "total_size": 27942091
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 860,
                                "total_size": 562950739
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 1127,
                                "total_size": 215448849
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 2070,
                                "total_size": 271112939
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 55,
                                "total_size": 27198937
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 93,
                                "total_size": 2998141
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 9238,
                                "total_size": 849950213
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 898,
                                "total_size": 705968969
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 511,
                                "total_size": 138374554
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 413,
                                "total_size": 157092655
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 6861,
                                "total_size": 648088876
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 81,
                                "total_size": 35554729
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 902,
                                "total_size": 91946068
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 398,
                                "total_size": 35724687
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 8,
                                "total_size": 235233448
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 7473,
                                "total_size": 635289466
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 547,
                                "total_size": 144834664
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 579,
                                "total_size": 1769275660
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 1270,
                                "total_size": 690661664
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 820,
                                "total_size": 76042821
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 279,
                                "total_size": 1701137216
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 6819,
                                "total_size": 662260033
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 730,
                                "total_size": 36484021
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 273,
                                "total_size": 1736350432
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 423,
                                "total_size": 97185401
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 168,
                                "total_size": 2565960
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 5338,
                                "total_size": 474703292
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 588,
                                "total_size": 124623889
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 29,
                                "total_size": 712010
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 38,
                                "total_size": 999967
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 211,
                                "total_size": 26139620
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 8105,
                                "total_size": 713448336
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 57,
                                "total_size": 48953760
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 44,
                                "total_size": 15292927
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 138,
                                "total_size": 13527142
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 96,
                                "total_size": 7097275
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 46,
                                "total_size": 5206787
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 7,
                                "total_size": 4283768
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 554,
                                "total_size": 120746276
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 948,
                                "total_size": 449408670
                            }
                        ]
                    },
                    "time": "2025-02-20T06:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "tJViJZUBV5X_QYDfAJwV",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 193433,
                        "total_size": 61865461004,
                        "group_by_agent": [
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 406,
                                "total_size": 213814288
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 22,
                                "total_size": 16489340
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 357,
                                "total_size": 238279561
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 29,
                                "total_size": 9760042
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 10,
                                "total_size": 308584
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 16,
                                "total_size": 1525427
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 433,
                                "total_size": 328360285
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 1188,
                                "total_size": 105407121
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 41,
                                "total_size": 490858
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 335,
                                "total_size": 122272191
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 45,
                                "total_size": 447282
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 302,
                                "total_size": 123800807
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 419,
                                "total_size": 332509058
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 1567,
                                "total_size": 140141452
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 278,
                                "total_size": 95004109
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 3805,
                                "total_size": 363236642
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 311,
                                "total_size": 32753582
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 1227,
                                "total_size": 834258377
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 24,
                                "total_size": 338608523
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 207,
                                "total_size": 24503954
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 361,
                                "total_size": 256558036
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 33,
                                "total_size": 19964140
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 5476,
                                "total_size": 530626055
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 247,
                                "total_size": 115230465
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 179,
                                "total_size": 54991693
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 219,
                                "total_size": 17642441
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 217,
                                "total_size": 56734633
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 262,
                                "total_size": 76665383
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 21,
                                "total_size": 534727
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 391,
                                "total_size": 148753299
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 4494,
                                "total_size": 424819777
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 1060,
                                "total_size": 1029632400
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 405,
                                "total_size": 37133125
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 1718,
                                "total_size": 116047524
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 441,
                                "total_size": 245001418
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 83,
                                "total_size": 174479268
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 490,
                                "total_size": 196577850
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 207,
                                "total_size": 17174522
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 360,
                                "total_size": 377752580
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 216,
                                "total_size": 48723463
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 396,
                                "total_size": 100074864
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 325,
                                "total_size": 77668326
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 415,
                                "total_size": 97130678
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 448,
                                "total_size": 85091851
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 12,
                                "total_size": 11127399
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 501,
                                "total_size": 1526185195
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 134,
                                "total_size": 74611986
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 1661,
                                "total_size": 168757385
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 127,
                                "total_size": 433066698
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 22,
                                "total_size": 2568605
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 304,
                                "total_size": 31168535
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 40,
                                "total_size": 1806296
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 340,
                                "total_size": 47857108
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 44,
                                "total_size": 681177444
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 77,
                                "total_size": 45197509
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 382,
                                "total_size": 24542973
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 22,
                                "total_size": 384970
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 20,
                                "total_size": 933855
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 8,
                                "total_size": 422110
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 685,
                                "total_size": 69864656
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 976,
                                "total_size": 461631382
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 5171,
                                "total_size": 487795764
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 919,
                                "total_size": 74456984
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 338,
                                "total_size": 500083231
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 177,
                                "total_size": 76108660
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 255,
                                "total_size": 243321984
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 4736,
                                "total_size": 418017460
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 759,
                                "total_size": 494620937
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 1100,
                                "total_size": 692227468
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 172,
                                "total_size": 25160269
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 259,
                                "total_size": 1609806282
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 1,
                                "total_size": 852986
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 350,
                                "total_size": 10071884
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 848,
                                "total_size": 223039078
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 29,
                                "total_size": 9190515
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 378,
                                "total_size": 34592388
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 497,
                                "total_size": 145112130
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 266,
                                "total_size": 174059111
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 192,
                                "total_size": 751661889
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 196,
                                "total_size": 77269411
                            },
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 194,
                                "total_size": 12208331
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 131,
                                "total_size": 32710407
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 237,
                                "total_size": 121832064
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 20,
                                "total_size": 1098468
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 412,
                                "total_size": 30814921
                            },
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 3712,
                                "total_size": 369185216
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 407,
                                "total_size": 101077280
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 255,
                                "total_size": 7309941
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 5,
                                "total_size": 261589
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 20,
                                "total_size": 1504027
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 4,
                                "total_size": 3324772
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 1674,
                                "total_size": 109342799
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 11,
                                "total_size": 189338
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 426,
                                "total_size": 243284411
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 1361,
                                "total_size": 1163147409
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 451,
                                "total_size": 189173318
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 110,
                                "total_size": 20722630
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 428,
                                "total_size": 141690186
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 233,
                                "total_size": 49341261
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 129,
                                "total_size": 227375676
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 394,
                                "total_size": 26538072
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 666,
                                "total_size": 225945893
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 287,
                                "total_size": 817487788
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 213,
                                "total_size": 3242613
                            },
                            {
                                "rc_agent": "0143a74d-9e52-4478-aba0-f0853a0d9f02",
                                "total_count": 393,
                                "total_size": 95747551
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 457,
                                "total_size": 114989487
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 34,
                                "total_size": 17071941
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 287,
                                "total_size": 115536960
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 220,
                                "total_size": 28230584
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 214,
                                "total_size": 25484955
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 966,
                                "total_size": 493258138
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 1025,
                                "total_size": 3144829577
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 425,
                                "total_size": 31046820
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 175,
                                "total_size": 88596325
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 340,
                                "total_size": 120744319
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 3331,
                                "total_size": 259805410
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 240,
                                "total_size": 52708088
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 525,
                                "total_size": 88339731
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 204,
                                "total_size": 126359667
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 349,
                                "total_size": 59813058
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 365,
                                "total_size": 110951618
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 385,
                                "total_size": 111915894
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 612,
                                "total_size": 48504467
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 145,
                                "total_size": 25437459
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 6364,
                                "total_size": 617744664
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 79,
                                "total_size": 69481461
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 527,
                                "total_size": 1299418611
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 130,
                                "total_size": 4919936
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 218,
                                "total_size": 24621194
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 817,
                                "total_size": 477518436
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 260,
                                "total_size": 26565935
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 258,
                                "total_size": 1647506292
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 187,
                                "total_size": 42567394
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 195,
                                "total_size": 44229406
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 843,
                                "total_size": 104045178
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 5049,
                                "total_size": 464432516
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 249,
                                "total_size": 266644146
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 157,
                                "total_size": 72617631
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 201,
                                "total_size": 164259117
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 222,
                                "total_size": 17479117
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 1379,
                                "total_size": 287362685
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 253,
                                "total_size": 57073919
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 15,
                                "total_size": 86302
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 227,
                                "total_size": 91683414
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 234,
                                "total_size": 38084506
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 350,
                                "total_size": 114939835
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 212,
                                "total_size": 156841985
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 317,
                                "total_size": 31595704
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 566,
                                "total_size": 435897439
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 634,
                                "total_size": 308136587
                            },
                            {
                                "rc_agent": "456d9bf9-a53a-4786-9c2c-24cdd5fd6771",
                                "total_count": 784,
                                "total_size": 341671021
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 2596,
                                "total_size": 223844970
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 2403,
                                "total_size": 204405132
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 2152,
                                "total_size": 181279716
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 450,
                                "total_size": 299417758
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 497,
                                "total_size": 1488377638
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 472,
                                "total_size": 43626891
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 5,
                                "total_size": 63928016
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 169,
                                "total_size": 37374258
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 81,
                                "total_size": 58180659
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 211,
                                "total_size": 29921315
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 3,
                                "total_size": 86995684
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 221,
                                "total_size": 43643152
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 274,
                                "total_size": 93768617
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 541,
                                "total_size": 139456167
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 413,
                                "total_size": 37218019
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 248,
                                "total_size": 1681529340
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 16,
                                "total_size": 2068486
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 329,
                                "total_size": 14850369
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 2643,
                                "total_size": 272366545
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 27,
                                "total_size": 100739484
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 228,
                                "total_size": 18914772
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 45,
                                "total_size": 1690509
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 23,
                                "total_size": 17780776
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 397,
                                "total_size": 366717063
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 248,
                                "total_size": 135321903
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 3414,
                                "total_size": 309333296
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 7,
                                "total_size": 53340
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 4431,
                                "total_size": 389725681
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 583,
                                "total_size": 147544605
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 210,
                                "total_size": 166851292
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 285,
                                "total_size": 148590411
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 473,
                                "total_size": 184766287
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 209,
                                "total_size": 21619370
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 165,
                                "total_size": 109000433
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 35,
                                "total_size": 891455
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 362,
                                "total_size": 84806802
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 232,
                                "total_size": 18514378
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 2283,
                                "total_size": 211298026
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 3221,
                                "total_size": 299825103
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 4227,
                                "total_size": 351338744
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 219,
                                "total_size": 17265685
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 37,
                                "total_size": 1625975
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 24,
                                "total_size": 7735377
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 5,
                                "total_size": 196979
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 315,
                                "total_size": 22636710
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 690,
                                "total_size": 123303400
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 28,
                                "total_size": 45805378
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 746,
                                "total_size": 73676784
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 1432,
                                "total_size": 433813540
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 396,
                                "total_size": 27516559
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 1882,
                                "total_size": 161862213
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 4385,
                                "total_size": 373577750
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 407,
                                "total_size": 281972824
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 709,
                                "total_size": 220659095
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 60,
                                "total_size": 1086055
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 23,
                                "total_size": 2610537
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 69,
                                "total_size": 2978708
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 380,
                                "total_size": 28833526
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 400,
                                "total_size": 153838664
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 450,
                                "total_size": 48123382
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 1541,
                                "total_size": 139067985
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 985,
                                "total_size": 2052620727
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 419,
                                "total_size": 31084199
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 979,
                                "total_size": 1796361262
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 9,
                                "total_size": 513453
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 3434,
                                "total_size": 216054640
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 205,
                                "total_size": 14973989
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 221,
                                "total_size": 16566603
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 520,
                                "total_size": 821145253
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 250,
                                "total_size": 11648016
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 277,
                                "total_size": 780453473
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 134,
                                "total_size": 36448305
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 12,
                                "total_size": 7315590
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 1524,
                                "total_size": 139984079
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 198,
                                "total_size": 107520001
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 165,
                                "total_size": 62080325
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 12,
                                "total_size": 183227954
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 31,
                                "total_size": 43017139
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 920,
                                "total_size": 445212744
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 364,
                                "total_size": 382041721
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 224,
                                "total_size": 56264850
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 24,
                                "total_size": 263455
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 112,
                                "total_size": 2049310
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 386,
                                "total_size": 47341726
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 302,
                                "total_size": 11832484
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 330,
                                "total_size": 10528977
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 2,
                                "total_size": 4260075
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 13,
                                "total_size": 20219685
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 325,
                                "total_size": 89840822
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 345,
                                "total_size": 88054295
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 378,
                                "total_size": 117406550
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 6,
                                "total_size": 553381
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 754,
                                "total_size": 76625002
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 352,
                                "total_size": 30666480
                            },
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 438,
                                "total_size": 110589759
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 266,
                                "total_size": 13622507
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 96,
                                "total_size": 9436001
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 843,
                                "total_size": 470249916
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 311,
                                "total_size": 141638909
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 5550,
                                "total_size": 440198987
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 4955,
                                "total_size": 462181978
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 462,
                                "total_size": 48792492
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 140,
                                "total_size": 87601138
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 74,
                                "total_size": 43312098
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 338,
                                "total_size": 75013345
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 258,
                                "total_size": 60625610
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 8,
                                "total_size": 11919413
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 339,
                                "total_size": 93316642
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 2823,
                                "total_size": 265499179
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 36,
                                "total_size": 1582089
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 571,
                                "total_size": 390708371
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 292,
                                "total_size": 128189306
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 124,
                                "total_size": 68307517
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 1377,
                                "total_size": 1041270381
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 86,
                                "total_size": 3478367
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 239,
                                "total_size": 16222743
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 98,
                                "total_size": 72924676
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 315,
                                "total_size": 126214717
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 6,
                                "total_size": 660295
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 1185,
                                "total_size": 842887753
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 1263,
                                "total_size": 1724322211
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 142,
                                "total_size": 7850240
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 206,
                                "total_size": 47276212
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 1501,
                                "total_size": 1085417899
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 136,
                                "total_size": 10180190
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 68,
                                "total_size": 1354957
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 611,
                                "total_size": 147136195
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 885,
                                "total_size": 89077790
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 2074,
                                "total_size": 91417522
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 261,
                                "total_size": 86096696
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 348,
                                "total_size": 9530566
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 1385,
                                "total_size": 799963798
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 17,
                                "total_size": 438975
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 2204,
                                "total_size": 226437483
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 282,
                                "total_size": 15325181
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 272,
                                "total_size": 12681398
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 263,
                                "total_size": 79680277
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 1060,
                                "total_size": 620215348
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 13,
                                "total_size": 514369
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 40,
                                "total_size": 834693
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 328,
                                "total_size": 12400023
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 218,
                                "total_size": 160208717
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 295,
                                "total_size": 81389743
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 283,
                                "total_size": 3753855
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 126,
                                "total_size": 44234896
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "YUYAJZUBT5H4H6HfCzau",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 67,
                        "total_size": 6384271,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 67,
                                "total_size": 6384271
                            }
                        ]
                    },
                    "time": "2025-02-20T09:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "aUYAJZUBT5H4H6HfCzau",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 70,
                        "total_size": 9616693,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 60,
                                "total_size": 5745539
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 10,
                                "total_size": 3871154
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "akYAJZUBT5H4H6HfCzau",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 79,
                        "total_size": 8189319,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 79,
                                "total_size": 8189319
                            }
                        ]
                    },
                    "time": "2025-02-20T08:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "6bykJZUBk0bmQQGusxiW",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T22:00:00.000000Z",
                    "indices": {
                        "total_count": 8,
                        "total_size": 4315565,
                        "group_by_agent": [
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 8,
                                "total_size": 4315565
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "vJXbJZUBV5X_QYDf4pxD",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T23:00:00.000000Z",
                    "indices": {
                        "total_count": 603,
                        "total_size": 49132687,
                        "group_by_agent": [
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 2,
                                "total_size": 169283
                            },
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 601,
                                "total_size": 48963404
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "vpXbJZUBV5X_QYDf4pxD",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T23:00:00.000000Z",
                    "indices": {
                        "total_count": 216,
                        "total_size": 17380561,
                        "group_by_agent": [
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 6,
                                "total_size": 431132
                            },
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 210,
                                "total_size": 16949429
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "cZVbJJUBLBoVSWi0Sx7K",
                "_score": 1.0041407,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T16:00:00.000000Z",
                    "indices": {
                        "total_count": 118,
                        "total_size": 15375428,
                        "group_by_agent": [
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 11,
                                "total_size": 2987342
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 98,
                                "total_size": 8301569
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 9,
                                "total_size": 4086517
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "k5ViJZUBV5X_QYDfAJwV",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 287901,
                        "total_size": 92113957982,
                        "group_by_agent": [
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 263,
                                "total_size": 113496713
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 640,
                                "total_size": 443716086
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 533,
                                "total_size": 28466227
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 6081,
                                "total_size": 580744290
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 7097,
                                "total_size": 651733490
                            },
                            {
                                "rc_agent": "cde365e1-868f-493d-9adf-7ce505eec802",
                                "total_count": 1,
                                "total_size": 17312
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 364,
                                "total_size": 90014070
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 28,
                                "total_size": 309125
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 437,
                                "total_size": 16617935
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 484,
                                "total_size": 60978923
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 3141,
                                "total_size": 270854035
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 322,
                                "total_size": 1963912157
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 1146,
                                "total_size": 116652420
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 5772,
                                "total_size": 519640902
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 829,
                                "total_size": 156593186
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 652,
                                "total_size": 153533849
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 43,
                                "total_size": 13404839
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 514,
                                "total_size": 118625904
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 178,
                                "total_size": 330573072
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 462,
                                "total_size": 444572629
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 3,
                                "total_size": 333450
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 49,
                                "total_size": 1019106
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 554,
                                "total_size": 505874757
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 344,
                                "total_size": 25769862
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 3189,
                                "total_size": 140344805
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 4,
                                "total_size": 171162
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 53,
                                "total_size": 2471692
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 301,
                                "total_size": 243674808
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 3524,
                                "total_size": 312229063
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 374,
                                "total_size": 371396386
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 950,
                                "total_size": 630720447
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 2080,
                                "total_size": 184574923
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 28,
                                "total_size": 26318004
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 613,
                                "total_size": 199517336
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 713,
                                "total_size": 428551968
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 126,
                                "total_size": 2511142
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 7,
                                "total_size": 196312358
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 1316,
                                "total_size": 734318985
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 319,
                                "total_size": 50419624
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 27,
                                "total_size": 243087829
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 30,
                                "total_size": 3610812
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 30,
                                "total_size": 761036
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 88,
                                "total_size": 1604132
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 6255,
                                "total_size": 529137255
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 211,
                                "total_size": 105837184
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 1965,
                                "total_size": 1624947374
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 272,
                                "total_size": 84444503
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 11,
                                "total_size": 1400029
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 449,
                                "total_size": 17410076
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 17,
                                "total_size": 608312
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 1278,
                                "total_size": 2190103780
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 14,
                                "total_size": 8493430
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 31,
                                "total_size": 23939151
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 628,
                                "total_size": 167064544
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 979,
                                "total_size": 248467102
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 603,
                                "total_size": 136865926
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 398,
                                "total_size": 137835178
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 37,
                                "total_size": 3023052
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 898,
                                "total_size": 374746542
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 566,
                                "total_size": 44269220
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 803,
                                "total_size": 461519870
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 42,
                                "total_size": 4858544
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 7792,
                                "total_size": 735545254
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 19,
                                "total_size": 327839
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 2,
                                "total_size": 140729
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 609,
                                "total_size": 246201134
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 2510,
                                "total_size": 169687206
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 3672,
                                "total_size": 339292498
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 3,
                                "total_size": 966436
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 523,
                                "total_size": 145728719
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 1605,
                                "total_size": 1104197042
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 6921,
                                "total_size": 628083389
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 505,
                                "total_size": 45187601
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 48,
                                "total_size": 563346
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 354,
                                "total_size": 32361397
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 728,
                                "total_size": 2223985670
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 6993,
                                "total_size": 551842536
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 282,
                                "total_size": 7735071
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 657,
                                "total_size": 484308110
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 551,
                                "total_size": 209478495
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 572,
                                "total_size": 399895979
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 2428,
                                "total_size": 151950312
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 2196,
                                "total_size": 466407772
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 4567,
                                "total_size": 433596631
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 4128,
                                "total_size": 379511063
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 2055,
                                "total_size": 630728765
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 351,
                                "total_size": 230119408
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 449,
                                "total_size": 104482360
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 934,
                                "total_size": 237760138
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 615,
                                "total_size": 101748003
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 607,
                                "total_size": 44624158
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 5000,
                                "total_size": 455226890
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 676,
                                "total_size": 50487793
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 1360,
                                "total_size": 180872399
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 688,
                                "total_size": 49660535
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 595,
                                "total_size": 77605511
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 94,
                                "total_size": 25554545
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 4881,
                                "total_size": 496165410
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 54,
                                "total_size": 14374984
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 1,
                                "total_size": 854797
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 133,
                                "total_size": 45432372
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 711,
                                "total_size": 219725487
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 4407,
                                "total_size": 366421525
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 461,
                                "total_size": 105019354
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 124,
                                "total_size": 114503089
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 20,
                                "total_size": 30113031
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 423,
                                "total_size": 12182784
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 30,
                                "total_size": 18212404
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 8,
                                "total_size": 916810
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 24,
                                "total_size": 628992
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 328,
                                "total_size": 183150498
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 136,
                                "total_size": 112908250
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 62,
                                "total_size": 26957277
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 302,
                                "total_size": 22216669
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 4200,
                                "total_size": 328168337
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 354,
                                "total_size": 48682797
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 39,
                                "total_size": 379437
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 916,
                                "total_size": 241954732
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 686,
                                "total_size": 55181277
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 374,
                                "total_size": 41869048
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 452,
                                "total_size": 437122689
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 515,
                                "total_size": 14857614
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 26,
                                "total_size": 20190346
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 200,
                                "total_size": 93907830
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 680,
                                "total_size": 230573249
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 20,
                                "total_size": 1045937
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 1627,
                                "total_size": 896248370
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 872,
                                "total_size": 40341352
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 583,
                                "total_size": 8022340
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 28,
                                "total_size": 9936247
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 473,
                                "total_size": 126898401
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 267,
                                "total_size": 108699216
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 364,
                                "total_size": 1241230536
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 3814,
                                "total_size": 323610530
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 929,
                                "total_size": 166670574
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 521,
                                "total_size": 70213587
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 215,
                                "total_size": 622165884
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 617,
                                "total_size": 246562079
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 796,
                                "total_size": 2067155004
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 129,
                                "total_size": 2322032
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 31,
                                "total_size": 179059
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 430,
                                "total_size": 49795801
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 732,
                                "total_size": 79605711
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 392,
                                "total_size": 27075037
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 494,
                                "total_size": 51173543
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 402,
                                "total_size": 37650638
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 149,
                                "total_size": 5998981
                            },
                            {
                                "rc_agent": "c18624e7-e7f5-48c1-9352-4965a0f8c59e",
                                "total_count": 1072,
                                "total_size": 93212408
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 456,
                                "total_size": 671236966
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 263,
                                "total_size": 11705142
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 577,
                                "total_size": 464751590
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 7,
                                "total_size": 1904923
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 86,
                                "total_size": 2711972
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 598,
                                "total_size": 135178973
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 294,
                                "total_size": 50405847
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 1567,
                                "total_size": 2828314833
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 938,
                                "total_size": 222708997
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 7832,
                                "total_size": 777856317
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 40,
                                "total_size": 99329644
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 934,
                                "total_size": 722203534
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 207,
                                "total_size": 136587401
                            },
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 2635,
                                "total_size": 175461730
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 2,
                                "total_size": 331037
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 5,
                                "total_size": 4023795
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 169,
                                "total_size": 7315461
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 278,
                                "total_size": 10400719
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 460,
                                "total_size": 125969694
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 51,
                                "total_size": 122806061
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 531,
                                "total_size": 125291839
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 484,
                                "total_size": 146796069
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 640,
                                "total_size": 161872980
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 405,
                                "total_size": 40482249
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 464,
                                "total_size": 99995356
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 323,
                                "total_size": 199460549
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 170,
                                "total_size": 122233103
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 4,
                                "total_size": 210340
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 492,
                                "total_size": 15670542
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 5810,
                                "total_size": 549922768
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 657,
                                "total_size": 32867971
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 138,
                                "total_size": 16405267
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 1207,
                                "total_size": 321921256
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 4,
                                "total_size": 67071952
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 32,
                                "total_size": 1393895
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 31,
                                "total_size": 81101364
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 530,
                                "total_size": 183695491
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 7,
                                "total_size": 403663
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 1463,
                                "total_size": 723904316
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 43,
                                "total_size": 13757845
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 424,
                                "total_size": 72744254
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 1529,
                                "total_size": 520053398
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 364,
                                "total_size": 2091505981
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 304,
                                "total_size": 157591774
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 1796,
                                "total_size": 1231315976
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 33,
                                "total_size": 51418092
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 37,
                                "total_size": 938097
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 347,
                                "total_size": 335144716
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 396,
                                "total_size": 40993815
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 2,
                                "total_size": 140939
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 83,
                                "total_size": 47082240
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 153,
                                "total_size": 9842942
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 349,
                                "total_size": 88520819
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 653,
                                "total_size": 58597430
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 481,
                                "total_size": 191919147
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 617,
                                "total_size": 320283535
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 1,
                                "total_size": 98480
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 1696,
                                "total_size": 987542793
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 13,
                                "total_size": 517343
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 2458,
                                "total_size": 1795042971
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 422,
                                "total_size": 159883375
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 1352,
                                "total_size": 859993412
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 661,
                                "total_size": 498542420
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 845,
                                "total_size": 198372085
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 888,
                                "total_size": 93828202
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 4880,
                                "total_size": 487272852
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 6118,
                                "total_size": 575693382
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 502,
                                "total_size": 175766659
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 580,
                                "total_size": 2411874391
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 118,
                                "total_size": 5429564
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 107,
                                "total_size": 6144397
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 311,
                                "total_size": 123599004
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 631,
                                "total_size": 47870561
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 246,
                                "total_size": 124667336
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 1535,
                                "total_size": 2983790935
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 882,
                                "total_size": 227258086
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 717,
                                "total_size": 216104179
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 13,
                                "total_size": 100270
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 1281,
                                "total_size": 527424795
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 123,
                                "total_size": 66997158
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 2706,
                                "total_size": 274743259
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 380,
                                "total_size": 102745609
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 388,
                                "total_size": 2341429426
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 42,
                                "total_size": 415843814
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 440,
                                "total_size": 188981810
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 4,
                                "total_size": 2233299
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 606,
                                "total_size": 44069216
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 133,
                                "total_size": 15032656
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 2203,
                                "total_size": 206329595
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 2637,
                                "total_size": 174154976
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 472,
                                "total_size": 29524021
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 764,
                                "total_size": 123119911
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 83,
                                "total_size": 222228992
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 689,
                                "total_size": 164342213
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 374,
                                "total_size": 209541191
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 3133,
                                "total_size": 277343623
                            },
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 3175,
                                "total_size": 316186059
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 43,
                                "total_size": 721091
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 1570,
                                "total_size": 1527804158
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 554,
                                "total_size": 370208841
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 43,
                                "total_size": 3197123
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 3215,
                                "total_size": 287445755
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 32,
                                "total_size": 18863678
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 1142,
                                "total_size": 112363954
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 6585,
                                "total_size": 534185794
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 512,
                                "total_size": 207621484
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 1447,
                                "total_size": 677301856
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 993,
                                "total_size": 100819626
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 1997,
                                "total_size": 1023204805
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 525,
                                "total_size": 55141683
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 777,
                                "total_size": 2385480543
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 401,
                                "total_size": 6156238
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 979,
                                "total_size": 291736891
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 38,
                                "total_size": 1829878
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 477,
                                "total_size": 119396193
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 309,
                                "total_size": 101960394
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 438,
                                "total_size": 204441477
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 605,
                                "total_size": 479304184
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 792,
                                "total_size": 607898842
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 600,
                                "total_size": 198722275
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 5374,
                                "total_size": 508051313
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 137,
                                "total_size": 26171191
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 593,
                                "total_size": 45459450
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 2,
                                "total_size": 595868
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 1331,
                                "total_size": 4119060429
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 17,
                                "total_size": 137299809
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 7290,
                                "total_size": 647296340
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 1938,
                                "total_size": 1378119925
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 36,
                                "total_size": 899158
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 21,
                                "total_size": 2122657
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 369,
                                "total_size": 33896754
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 615,
                                "total_size": 182563441
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 670,
                                "total_size": 262031027
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 44,
                                "total_size": 2020370
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 229,
                                "total_size": 143682692
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 1006,
                                "total_size": 87140184
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 443,
                                "total_size": 1250176641
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 1907,
                                "total_size": 1238789667
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 89,
                                "total_size": 8712638
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 1805,
                                "total_size": 2469194655
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 1141,
                                "total_size": 830075071
                            }
                        ]
                    },
                    "time": "2025-02-20T08:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "r5ViJZUBV5X_QYDfAJwV",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 448942,
                        "total_size": 133034708644,
                        "group_by_agent": [
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 189,
                                "total_size": 18531641
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 852,
                                "total_size": 64659563
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 745,
                                "total_size": 494931836
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 41,
                                "total_size": 31508053
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 99,
                                "total_size": 980993
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 2078,
                                "total_size": 1486568999
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 683,
                                "total_size": 1938925914
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 1074,
                                "total_size": 75652625
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 108,
                                "total_size": 3386710
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 95,
                                "total_size": 45693172
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 5524,
                                "total_size": 489797186
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 840,
                                "total_size": 685129892
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 31,
                                "total_size": 1610466
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 565,
                                "total_size": 8553120
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 487,
                                "total_size": 272536821
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 384,
                                "total_size": 29970458
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 539,
                                "total_size": 135802620
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 1276,
                                "total_size": 945862333
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 1104,
                                "total_size": 332527458
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 625,
                                "total_size": 150097938
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 707,
                                "total_size": 22432046
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 2158,
                                "total_size": 1051497831
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 11661,
                                "total_size": 989687884
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 21,
                                "total_size": 186190242
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 1300,
                                "total_size": 2493162861
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 510,
                                "total_size": 209515239
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 1797,
                                "total_size": 982289921
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 11383,
                                "total_size": 1046582643
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 792,
                                "total_size": 708946580
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 827,
                                "total_size": 653301524
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 516,
                                "total_size": 291962247
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 27,
                                "total_size": 206873
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 544,
                                "total_size": 61264004
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 15,
                                "total_size": 585055
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 483,
                                "total_size": 21610468
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 1446,
                                "total_size": 273748270
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 2066,
                                "total_size": 1399682828
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 11052,
                                "total_size": 1038867390
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 11388,
                                "total_size": 1046609942
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 60,
                                "total_size": 15072562
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 364,
                                "total_size": 529352313
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 5886,
                                "total_size": 4909999095
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 682,
                                "total_size": 525780329
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 608,
                                "total_size": 202510643
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 633,
                                "total_size": 188811319
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 41,
                                "total_size": 32580815
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 535,
                                "total_size": 72150342
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 1341,
                                "total_size": 350423662
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 207,
                                "total_size": 99936771
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 149,
                                "total_size": 115666690
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 41,
                                "total_size": 68954999
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 527,
                                "total_size": 38665558
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 6036,
                                "total_size": 1803786069
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 979,
                                "total_size": 70739043
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 657,
                                "total_size": 600722596
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 1117,
                                "total_size": 265614640
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 1554,
                                "total_size": 993710812
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 2100,
                                "total_size": 6059324834
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 172,
                                "total_size": 16865728
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 619,
                                "total_size": 83859174
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 5674,
                                "total_size": 543595912
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 1125,
                                "total_size": 310825620
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 1086,
                                "total_size": 250622265
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 554,
                                "total_size": 50741862
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 132,
                                "total_size": 96477497
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 334,
                                "total_size": 2099436341
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 1052,
                                "total_size": 355549236
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 10,
                                "total_size": 1892853
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 3616,
                                "total_size": 244266975
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 1273,
                                "total_size": 312804296
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 44,
                                "total_size": 12871702
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 1668,
                                "total_size": 282692129
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 881,
                                "total_size": 2562870747
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 281,
                                "total_size": 509451525
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 459,
                                "total_size": 306628772
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 107,
                                "total_size": 51363503
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 583,
                                "total_size": 27348557
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 1048,
                                "total_size": 77810455
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 314,
                                "total_size": 193580503
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 677,
                                "total_size": 239583367
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 285,
                                "total_size": 10730325
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 607,
                                "total_size": 54915673
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 828,
                                "total_size": 323480996
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 857,
                                "total_size": 435628256
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 426,
                                "total_size": 1737638546
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 2570,
                                "total_size": 1855461163
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 2018,
                                "total_size": 1169103135
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 1159,
                                "total_size": 804872633
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 502,
                                "total_size": 3422701183
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 219,
                                "total_size": 68072347
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 39,
                                "total_size": 1162925
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 14684,
                                "total_size": 1323528654
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 850,
                                "total_size": 193309333
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 651,
                                "total_size": 182004677
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 36,
                                "total_size": 2644552
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 1739,
                                "total_size": 811293516
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 10331,
                                "total_size": 1000164244
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 3278,
                                "total_size": 871591758
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 419,
                                "total_size": 44049835
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 217,
                                "total_size": 14625312
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 3026,
                                "total_size": 1035077687
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 7538,
                                "total_size": 702814606
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 1190,
                                "total_size": 680950598
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 3226,
                                "total_size": 2222941479
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 105,
                                "total_size": 974753369
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 29,
                                "total_size": 2766179
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 77,
                                "total_size": 737509839
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 149,
                                "total_size": 6718147
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 44,
                                "total_size": 1145160
                            },
                            {
                                "rc_agent": "cde365e1-868f-493d-9adf-7ce505eec802",
                                "total_count": 1,
                                "total_size": 82704
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 1270,
                                "total_size": 325194730
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 450,
                                "total_size": 121099580
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 2472,
                                "total_size": 5138608126
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 625,
                                "total_size": 608009466
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 920,
                                "total_size": 122586170
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 597,
                                "total_size": 243708729
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 1182,
                                "total_size": 958624934
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 585,
                                "total_size": 145615057
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 12,
                                "total_size": 11726756
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 244,
                                "total_size": 66204589
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 2925,
                                "total_size": 394185935
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 1253,
                                "total_size": 283564258
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 1267,
                                "total_size": 130434009
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 299,
                                "total_size": 141019716
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 434,
                                "total_size": 146632124
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 506,
                                "total_size": 6708876
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 462,
                                "total_size": 56238005
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 904,
                                "total_size": 69080346
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 85,
                                "total_size": 51542773
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 37,
                                "total_size": 22507628
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 10,
                                "total_size": 175800143
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 1090,
                                "total_size": 97885074
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 9854,
                                "total_size": 867100013
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 77,
                                "total_size": 1954499
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 1368,
                                "total_size": 579181895
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 98,
                                "total_size": 35099418
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 22,
                                "total_size": 861041
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 11,
                                "total_size": 119657202
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 1127,
                                "total_size": 112255627
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 921,
                                "total_size": 1606844503
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 284,
                                "total_size": 148412460
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 384,
                                "total_size": 379527968
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 245,
                                "total_size": 10676322
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 178,
                                "total_size": 3280043
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 46,
                                "total_size": 5281739
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 198,
                                "total_size": 118057381
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 837,
                                "total_size": 203795419
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 50,
                                "total_size": 386479780
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 709,
                                "total_size": 420679375
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 611,
                                "total_size": 3901531309
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 910,
                                "total_size": 218079287
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 770,
                                "total_size": 306719462
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 3,
                                "total_size": 1552907
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 869,
                                "total_size": 66961983
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 1029,
                                "total_size": 106117841
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 2088,
                                "total_size": 1941552300
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 590,
                                "total_size": 59441979
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 292,
                                "total_size": 191003061
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 730,
                                "total_size": 56319911
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 394,
                                "total_size": 376056171
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 20,
                                "total_size": 344013
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 207,
                                "total_size": 3978174
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 44,
                                "total_size": 252526
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 754,
                                "total_size": 164759460
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 2084,
                                "total_size": 1195545537
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 9844,
                                "total_size": 908920407
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 366,
                                "total_size": 126856008
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 15,
                                "total_size": 837479
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 55,
                                "total_size": 2501437
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 763,
                                "total_size": 623338793
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 4741,
                                "total_size": 390251648
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 452,
                                "total_size": 38464148
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 683,
                                "total_size": 347137097
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 1256,
                                "total_size": 501719583
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 4237,
                                "total_size": 430853714
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 89,
                                "total_size": 157964886
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 12,
                                "total_size": 382002770
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 13172,
                                "total_size": 1290826136
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 1853,
                                "total_size": 185560316
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 3451,
                                "total_size": 4555259032
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 935,
                                "total_size": 46995808
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 7761,
                                "total_size": 670311304
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 421,
                                "total_size": 69905654
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 1034,
                                "total_size": 105943189
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 9734,
                                "total_size": 885413053
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 663,
                                "total_size": 18322149
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 84,
                                "total_size": 1417377
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 844,
                                "total_size": 2055903242
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 181,
                                "total_size": 41557338
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 5730,
                                "total_size": 555459195
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 1381,
                                "total_size": 147903804
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 35,
                                "total_size": 8187501
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 701,
                                "total_size": 266171015
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 870,
                                "total_size": 229589023
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 1186,
                                "total_size": 358437057
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 36,
                                "total_size": 11559627
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 72,
                                "total_size": 19557443
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 10695,
                                "total_size": 972803675
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 1672,
                                "total_size": 269577153
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 125,
                                "total_size": 228553738
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 184,
                                "total_size": 177892455
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 48,
                                "total_size": 520064
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 43,
                                "total_size": 5034870
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 63,
                                "total_size": 15041521
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 19,
                                "total_size": 2446098
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 1421,
                                "total_size": 429443720
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 10248,
                                "total_size": 809115979
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 51,
                                "total_size": 2127705
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 549,
                                "total_size": 39065985
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 4,
                                "total_size": 3123656
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 4790,
                                "total_size": 491305200
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 577,
                                "total_size": 55747704
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 236,
                                "total_size": 13256773
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 884,
                                "total_size": 2615277509
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 118,
                                "total_size": 70093300
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 680,
                                "total_size": 53054127
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 1068,
                                "total_size": 252304430
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 1122,
                                "total_size": 329787844
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 29,
                                "total_size": 43776466
                            },
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 119,
                                "total_size": 29182220
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 167,
                                "total_size": 3121247
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 134,
                                "total_size": 73980700
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 100,
                                "total_size": 4571628
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 661,
                                "total_size": 346660576
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 378,
                                "total_size": 42798151
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 17376,
                                "total_size": 1695667245
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 1057,
                                "total_size": 110573421
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 1416,
                                "total_size": 360299479
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 8047,
                                "total_size": 762948581
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 4133,
                                "total_size": 366598679
                            },
                            {
                                "rc_agent": "bace3ddf-e238-4660-9232-4d661fe426bf",
                                "total_count": 1,
                                "total_size": 30780
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 874,
                                "total_size": 341131620
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 67,
                                "total_size": 83870821
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 152,
                                "total_size": 10679449
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 473,
                                "total_size": 80150404
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 599,
                                "total_size": 17369169
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 992,
                                "total_size": 38858109
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 477,
                                "total_size": 296983561
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 571,
                                "total_size": 16434969
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 119,
                                "total_size": 8756043
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 1046,
                                "total_size": 269429540
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 1395,
                                "total_size": 558675844
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 1214,
                                "total_size": 937992564
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 645,
                                "total_size": 34318381
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 130,
                                "total_size": 35288984
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 170,
                                "total_size": 13988714
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 661,
                                "total_size": 431702397
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 561,
                                "total_size": 226256073
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 8,
                                "total_size": 427582
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 7,
                                "total_size": 34580295
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 490,
                                "total_size": 18999476
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 8,
                                "total_size": 916472
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 51,
                                "total_size": 3609165
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 4317,
                                "total_size": 386732667
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 6513,
                                "total_size": 580054754
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 5961,
                                "total_size": 394347112
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 2342,
                                "total_size": 492137693
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 232,
                                "total_size": 99357642
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 55,
                                "total_size": 1408914
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 721,
                                "total_size": 236439574
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 101,
                                "total_size": 31265607
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 95,
                                "total_size": 1980968
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 811,
                                "total_size": 307248256
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 1035,
                                "total_size": 106433451
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 1068,
                                "total_size": 83328266
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 43,
                                "total_size": 1708016
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 2038,
                                "total_size": 827882651
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 835,
                                "total_size": 73381873
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 145,
                                "total_size": 44935603
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 780,
                                "total_size": 171388712
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 1102,
                                "total_size": 82063320
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 39,
                                "total_size": 13163483
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 1299,
                                "total_size": 611114217
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 74,
                                "total_size": 864076
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 1035,
                                "total_size": 718977182
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 1544,
                                "total_size": 277332179
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 236,
                                "total_size": 77800264
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 730,
                                "total_size": 2046892314
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 378,
                                "total_size": 44388455
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 160,
                                "total_size": 30600395
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 1426,
                                "total_size": 62923140
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 889,
                                "total_size": 55991459
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 216,
                                "total_size": 73196615
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 1248,
                                "total_size": 779016444
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 708,
                                "total_size": 64584726
                            },
                            {
                                "rc_agent": "e5bfa1ac-6282-4695-b99f-b948c9244adc",
                                "total_count": 2007,
                                "total_size": 179803332
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 8359,
                                "total_size": 520711706
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 9182,
                                "total_size": 865455475
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 255,
                                "total_size": 41200643
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 795,
                                "total_size": 568376271
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 260,
                                "total_size": 886593240
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 6552,
                                "total_size": 590029840
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 3381,
                                "total_size": 291852034
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 643,
                                "total_size": 52867783
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 3471,
                                "total_size": 1744813720
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 279,
                                "total_size": 70454999
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 209,
                                "total_size": 8389811
                            }
                        ]
                    },
                    "time": "2025-02-20T05:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "tZViJZUBV5X_QYDfAJwV",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 217868,
                        "total_size": 70531895016,
                        "group_by_agent": [
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 21,
                                "total_size": 373509
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 4848,
                                "total_size": 382752807
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 1405,
                                "total_size": 1167522408
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 508,
                                "total_size": 154297901
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 247,
                                "total_size": 92619257
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 393,
                                "total_size": 90356295
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 336,
                                "total_size": 30174038
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 11,
                                "total_size": 9025263
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 1628,
                                "total_size": 3250009203
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 223,
                                "total_size": 91285201
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 3612,
                                "total_size": 336090041
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 265,
                                "total_size": 44073252
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 1391,
                                "total_size": 809719064
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 33,
                                "total_size": 2483584
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 2,
                                "total_size": 77730
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 116,
                                "total_size": 67785648
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 633,
                                "total_size": 65955931
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 543,
                                "total_size": 57806605
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 233,
                                "total_size": 54656800
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 501,
                                "total_size": 507289158
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 244,
                                "total_size": 105540947
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 18,
                                "total_size": 14247306
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 908,
                                "total_size": 76582603
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 5290,
                                "total_size": 498088307
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 135,
                                "total_size": 78099102
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 650,
                                "total_size": 2021693357
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 2530,
                                "total_size": 217948758
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 473,
                                "total_size": 78673939
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 20,
                                "total_size": 244181
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 410,
                                "total_size": 273271417
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 11,
                                "total_size": 16747820
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 137,
                                "total_size": 5120829
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 135,
                                "total_size": 9347637
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 16,
                                "total_size": 9668993
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 498,
                                "total_size": 730785479
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 463,
                                "total_size": 107073054
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 307,
                                "total_size": 1046862018
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 24,
                                "total_size": 18374985
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 346,
                                "total_size": 75351347
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 5883,
                                "total_size": 531503940
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 279,
                                "total_size": 228868758
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 12,
                                "total_size": 621396
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 265,
                                "total_size": 36679687
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 380,
                                "total_size": 114886851
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 466,
                                "total_size": 139030983
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 473,
                                "total_size": 35818825
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 492,
                                "total_size": 488627192
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 130,
                                "total_size": 15506054
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 2249,
                                "total_size": 176626505
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 775,
                                "total_size": 127578810
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 1312,
                                "total_size": 1226878760
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 229,
                                "total_size": 11844260
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 487,
                                "total_size": 50249137
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 5320,
                                "total_size": 488628664
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 151,
                                "total_size": 4169044
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 25,
                                "total_size": 1134079
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 55,
                                "total_size": 33452729
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 490,
                                "total_size": 171651068
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 26,
                                "total_size": 71281814
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 9,
                                "total_size": 406665
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 474,
                                "total_size": 451340408
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 49,
                                "total_size": 477474
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 138,
                                "total_size": 101115253
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 541,
                                "total_size": 373643029
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 639,
                                "total_size": 342426049
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 140,
                                "total_size": 14547569
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 458,
                                "total_size": 300399310
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 17,
                                "total_size": 5774510
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 282,
                                "total_size": 145470783
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 399,
                                "total_size": 93022360
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 44,
                                "total_size": 8657126
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 6,
                                "total_size": 159472
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 2,
                                "total_size": 1517941
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 1050,
                                "total_size": 107076871
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 582,
                                "total_size": 426354111
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 56,
                                "total_size": 85400273
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 567,
                                "total_size": 370166607
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 8,
                                "total_size": 44571856
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 460,
                                "total_size": 35949694
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 126,
                                "total_size": 28659499
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 68,
                                "total_size": 52104932
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 354,
                                "total_size": 19559414
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 23,
                                "total_size": 310156573
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 2678,
                                "total_size": 231044026
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 1076,
                                "total_size": 674035759
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 31,
                                "total_size": 73536238
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 501,
                                "total_size": 190436429
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 601,
                                "total_size": 1457481670
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 653,
                                "total_size": 117410923
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 36,
                                "total_size": 1094308
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 517,
                                "total_size": 177941878
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 443,
                                "total_size": 39171683
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 525,
                                "total_size": 286916679
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 328,
                                "total_size": 994587018
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 564,
                                "total_size": 221020037
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 408,
                                "total_size": 6186694
                            },
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 789,
                                "total_size": 51559826
                            },
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 3858,
                                "total_size": 384410803
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 1001,
                                "total_size": 339886552
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 548,
                                "total_size": 39943887
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 6465,
                                "total_size": 624360307
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 259,
                                "total_size": 202056832
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 221,
                                "total_size": 10271147
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 1,
                                "total_size": 109948
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 3399,
                                "total_size": 309297663
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 468,
                                "total_size": 113611782
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 1673,
                                "total_size": 1211255187
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 204,
                                "total_size": 126103005
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 195,
                                "total_size": 99102472
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 136,
                                "total_size": 5933119
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 415,
                                "total_size": 218148614
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 458,
                                "total_size": 1888842295
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 4629,
                                "total_size": 304563783
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 846,
                                "total_size": 227275873
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 82,
                                "total_size": 1618094
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 487,
                                "total_size": 509874099
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 403,
                                "total_size": 53220954
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 623,
                                "total_size": 346553253
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 380,
                                "total_size": 155408326
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 355,
                                "total_size": 96792988
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 45,
                                "total_size": 493351
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 226,
                                "total_size": 1365048248
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 1644,
                                "total_size": 827614179
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 276,
                                "total_size": 74404446
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 582,
                                "total_size": 43438437
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 59,
                                "total_size": 3483977
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 90,
                                "total_size": 8789984
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 268,
                                "total_size": 208966243
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 772,
                                "total_size": 359753049
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 315,
                                "total_size": 75682475
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 6943,
                                "total_size": 672907856
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 5677,
                                "total_size": 470993355
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 88,
                                "total_size": 15069447
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 338,
                                "total_size": 134738278
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 1087,
                                "total_size": 110581421
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 105,
                                "total_size": 3323017
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 2959,
                                "total_size": 272305377
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 29,
                                "total_size": 83985075
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 710,
                                "total_size": 182120848
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 126,
                                "total_size": 58678017
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 22,
                                "total_size": 713396
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 301,
                                "total_size": 33399406
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 1399,
                                "total_size": 938314440
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 7,
                                "total_size": 40941
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 295,
                                "total_size": 21319285
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 5770,
                                "total_size": 546510660
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 322,
                                "total_size": 84355680
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 231,
                                "total_size": 59225601
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 6,
                                "total_size": 752116
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 418,
                                "total_size": 2667685486
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 437,
                                "total_size": 187667011
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 3600,
                                "total_size": 321736732
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 3282,
                                "total_size": 289835534
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 1137,
                                "total_size": 111401560
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 1466,
                                "total_size": 308895741
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 121,
                                "total_size": 27439134
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 321,
                                "total_size": 29362283
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 43,
                                "total_size": 14931784
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 137,
                                "total_size": 99026399
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 188,
                                "total_size": 117306831
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 551,
                                "total_size": 143150944
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 168,
                                "total_size": 15372119
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 347,
                                "total_size": 2137965358
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 641,
                                "total_size": 1868358630
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 257,
                                "total_size": 7349556
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 4809,
                                "total_size": 454605420
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 107,
                                "total_size": 11091491
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 930,
                                "total_size": 1708469095
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 42,
                                "total_size": 1076105
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 955,
                                "total_size": 653664523
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 426,
                                "total_size": 38519042
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 561,
                                "total_size": 60551333
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 60,
                                "total_size": 1242352
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 5034,
                                "total_size": 442645481
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 1234,
                                "total_size": 3537930334
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 1163,
                                "total_size": 154812689
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 373,
                                "total_size": 102364745
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 2409,
                                "total_size": 246397277
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 515,
                                "total_size": 130053337
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 421,
                                "total_size": 1203321784
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 26,
                                "total_size": 482838
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 320,
                                "total_size": 4567838
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 17,
                                "total_size": 774301
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 4,
                                "total_size": 26033330
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 265,
                                "total_size": 24210821
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 250,
                                "total_size": 7208215
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 567,
                                "total_size": 989366417
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 97,
                                "total_size": 4448003
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 529,
                                "total_size": 215536526
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 147,
                                "total_size": 16715785
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 4113,
                                "total_size": 373178798
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 39,
                                "total_size": 308846293
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 377,
                                "total_size": 141125611
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 87,
                                "total_size": 3511091
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 405,
                                "total_size": 350361650
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 220,
                                "total_size": 104067843
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 269,
                                "total_size": 18598585
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 548,
                                "total_size": 49290891
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 943,
                                "total_size": 90014100
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 1785,
                                "total_size": 112165695
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 76,
                                "total_size": 44485495
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 582,
                                "total_size": 148045037
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 21,
                                "total_size": 6662278
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 301,
                                "total_size": 79898068
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 12,
                                "total_size": 207657
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 406,
                                "total_size": 264594234
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 4,
                                "total_size": 223136
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 645,
                                "total_size": 363611996
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 616,
                                "total_size": 142729877
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 385,
                                "total_size": 14973646
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 5,
                                "total_size": 464108
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 3,
                                "total_size": 22468
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 284,
                                "total_size": 69019433
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 14,
                                "total_size": 241251
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 330,
                                "total_size": 21542708
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 1241,
                                "total_size": 824015787
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 172,
                                "total_size": 97603135
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 787,
                                "total_size": 322430975
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 260,
                                "total_size": 101456412
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 405,
                                "total_size": 162669758
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 1,
                                "total_size": 28002
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 414,
                                "total_size": 27062568
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 650,
                                "total_size": 198290867
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 4,
                                "total_size": 127825552
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 901,
                                "total_size": 447944243
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 393,
                                "total_size": 131022807
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 86,
                                "total_size": 138195845
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 445,
                                "total_size": 146356449
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 11,
                                "total_size": 369284
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 275,
                                "total_size": 59794642
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 1337,
                                "total_size": 406953950
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 202,
                                "total_size": 112746223
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 4781,
                                "total_size": 482485545
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 480,
                                "total_size": 35143668
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 261,
                                "total_size": 59375099
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 2,
                                "total_size": 103861
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 231,
                                "total_size": 31268515
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 3438,
                                "total_size": 279590003
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 1644,
                                "total_size": 2265967975
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 602,
                                "total_size": 400309415
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 102,
                                "total_size": 92176768
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 4,
                                "total_size": 157829
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 3552,
                                "total_size": 156587047
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 259,
                                "total_size": 24908308
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 3,
                                "total_size": 4258992
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 609,
                                "total_size": 43887042
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 22,
                                "total_size": 7108198
                            },
                            {
                                "rc_agent": "cde365e1-868f-493d-9adf-7ce505eec802",
                                "total_count": 2,
                                "total_size": 627696
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 155,
                                "total_size": 5645929
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 34,
                                "total_size": 2122682
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 205,
                                "total_size": 135257120
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 434,
                                "total_size": 130437602
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 626,
                                "total_size": 117572973
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 11,
                                "total_size": 1282652
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 298,
                                "total_size": 115446658
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 526,
                                "total_size": 71041146
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 92,
                                "total_size": 25022914
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 1356,
                                "total_size": 972158048
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 399,
                                "total_size": 30820977
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 485,
                                "total_size": 130678697
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 422,
                                "total_size": 43526077
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 215,
                                "total_size": 117315642
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 3085,
                                "total_size": 208493950
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 4863,
                                "total_size": 430332741
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 554,
                                "total_size": 25768684
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 404,
                                "total_size": 137897383
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 221,
                                "total_size": 114529427
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 5302,
                                "total_size": 448018347
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 4838,
                                "total_size": 456057401
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 143,
                                "total_size": 112692114
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 234,
                                "total_size": 71133528
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 26,
                                "total_size": 2948667
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 16,
                                "total_size": 25370199
                            }
                        ]
                    },
                    "time": "2025-02-20T09:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "bJVbJJUBLBoVSWi0Sx7K",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T16:00:00.000000Z",
                    "indices": {
                        "total_count": 330,
                        "total_size": 45215331,
                        "group_by_agent": [
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 92,
                                "total_size": 7880232
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 109,
                                "total_size": 7319957
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 5,
                                "total_size": 1356776
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 6,
                                "total_size": 818612
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 18,
                                "total_size": 9499043
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 16,
                                "total_size": 7393090
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 74,
                                "total_size": 6384961
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 10,
                                "total_size": 4562660
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "cpVbJJUBLBoVSWi0Sx7K",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T16:00:00.000000Z",
                    "indices": {
                        "total_count": 5,
                        "total_size": 1358343,
                        "group_by_agent": [
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 5,
                                "total_size": 1358343
                            }
                        ]
                    },
                    "time": "2025-02-20T09:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "Mtw2JZUBWlvdGmPA9qZO",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T20:00:00.000000Z",
                    "indices": {
                        "total_count": 16,
                        "total_size": 6233460,
                        "group_by_agent": [
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 16,
                                "total_size": 6233460
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "aEYAJZUBT5H4H6HfCzau",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 56,
                        "total_size": 7717126,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 48,
                                "total_size": 4653933
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 8,
                                "total_size": 3063193
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "SNxtJZUBWlvdGmPA86Y9",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T21:00:00.000000Z",
                    "indices": {
                        "total_count": 53,
                        "total_size": 26084023,
                        "group_by_agent": [
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 41,
                                "total_size": 22823947
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 12,
                                "total_size": 3260076
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "0rySJJUBk0bmQQGuSRjp",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T17:00:00.000000Z",
                    "indices": {
                        "total_count": 77,
                        "total_size": 6359084,
                        "group_by_agent": [
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 77,
                                "total_size": 6359084
                            }
                        ]
                    },
                    "time": "2025-02-20T06:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "1bySJJUBk0bmQQGuSRjp",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T17:00:00.000000Z",
                    "indices": {
                        "total_count": 33,
                        "total_size": 2606307,
                        "group_by_agent": [
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 33,
                                "total_size": 2606307
                            }
                        ]
                    },
                    "time": "2025-02-20T10:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "1rySJJUBk0bmQQGuSRjp",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T17:00:00.000000Z",
                    "indices": {
                        "total_count": 86,
                        "total_size": 7045375,
                        "group_by_agent": [
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 86,
                                "total_size": 7045375
                            }
                        ]
                    },
                    "time": "2025-02-20T09:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "17ySJJUBk0bmQQGuSRjp",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T17:00:00.000000Z",
                    "indices": {
                        "total_count": 2,
                        "total_size": 174650,
                        "group_by_agent": [
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 2,
                                "total_size": 174650
                            }
                        ]
                    },
                    "time": "2025-02-20T07:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "dpXJJJUBLBoVSWi0Gx7k",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T18:00:00.000000Z",
                    "indices": {
                        "total_count": 405,
                        "total_size": 30987658,
                        "group_by_agent": [
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 153,
                                "total_size": 14167545
                            },
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 252,
                                "total_size": 16820113
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "Y9xJJpUBWlvdGmPA0qZ8",
                "_score": 1.0039448,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-21T01:00:00.000000Z",
                    "indices": {
                        "total_count": 31,
                        "total_size": 2822747,
                        "group_by_agent": [
                            {
                                "rc_agent": "c18624e7-e7f5-48c1-9352-4965a0f8c59e",
                                "total_count": 31,
                                "total_size": 2822747
                            }
                        ]
                    },
                    "time": "2025-02-20T07:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "lJViJZUBV5X_QYDfAJwV",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 317129,
                        "total_size": 102572713598,
                        "group_by_agent": [
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 5922,
                                "total_size": 589283460
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 1910,
                                "total_size": 1107928126
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 599,
                                "total_size": 60733629
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 654,
                                "total_size": 59681256
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 1,
                                "total_size": 858614
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 31,
                                "total_size": 788300
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 1181,
                                "total_size": 553574649
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 2979,
                                "total_size": 282074974
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 10461,
                                "total_size": 961766353
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 2191,
                                "total_size": 1837804865
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 640,
                                "total_size": 49995256
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 42,
                                "total_size": 1258728
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 932,
                                "total_size": 769579894
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 8,
                                "total_size": 2083945
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 185,
                                "total_size": 8350065
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 656,
                                "total_size": 479075883
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 341,
                                "total_size": 39849841
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 18,
                                "total_size": 27071752
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 4932,
                                "total_size": 454178139
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 46,
                                "total_size": 3912432
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 38,
                                "total_size": 55294028
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 1538,
                                "total_size": 4763375204
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 1549,
                                "total_size": 3199769288
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 182,
                                "total_size": 31197805
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 6,
                                "total_size": 25846764
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 99,
                                "total_size": 26611052
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 356,
                                "total_size": 85426413
                            },
                            {
                                "rc_agent": "cde365e1-868f-493d-9adf-7ce505eec802",
                                "total_count": 1,
                                "total_size": 528
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 2083,
                                "total_size": 199309818
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 657,
                                "total_size": 50742917
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 507,
                                "total_size": 478041122
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 528,
                                "total_size": 69353383
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 936,
                                "total_size": 440244496
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 1081,
                                "total_size": 258328149
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 28,
                                "total_size": 199558690
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 323,
                                "total_size": 44173503
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 563,
                                "total_size": 289823044
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 9964,
                                "total_size": 976272193
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 2494,
                                "total_size": 743584546
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 287,
                                "total_size": 89044734
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 655,
                                "total_size": 153293708
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 17,
                                "total_size": 2177707
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 980,
                                "total_size": 488489837
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 374,
                                "total_size": 27900399
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 1167,
                                "total_size": 628448350
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 718,
                                "total_size": 280175601
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 1116,
                                "total_size": 180933993
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 432,
                                "total_size": 1250144466
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 1046,
                                "total_size": 695608014
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 416,
                                "total_size": 2786358468
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 780,
                                "total_size": 318189411
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 74,
                                "total_size": 5402603
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 87,
                                "total_size": 2012237
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 7052,
                                "total_size": 572454962
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 907,
                                "total_size": 230041792
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 960,
                                "total_size": 218109485
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 1884,
                                "total_size": 3554489328
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 414,
                                "total_size": 232104747
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 39,
                                "total_size": 1807771
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 264,
                                "total_size": 139801251
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 132,
                                "total_size": 72400934
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 268,
                                "total_size": 99340839
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 1557,
                                "total_size": 148865594
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 957,
                                "total_size": 705240795
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 1047,
                                "total_size": 200295453
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 648,
                                "total_size": 520533921
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 177,
                                "total_size": 8311296
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 857,
                                "total_size": 525970937
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 180,
                                "total_size": 33864700
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 5251,
                                "total_size": 464262720
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 560,
                                "total_size": 43662707
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 30,
                                "total_size": 26896595
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 669,
                                "total_size": 1845323835
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 8205,
                                "total_size": 792181117
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 488,
                                "total_size": 26440707
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 25,
                                "total_size": 39592317
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 3,
                                "total_size": 163630
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 1233,
                                "total_size": 700870265
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 304,
                                "total_size": 26172017
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 620,
                                "total_size": 414397526
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 655,
                                "total_size": 679706118
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 628,
                                "total_size": 211408430
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 1179,
                                "total_size": 354550409
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 1870,
                                "total_size": 155241020
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 33,
                                "total_size": 835663
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 759,
                                "total_size": 299388670
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 314,
                                "total_size": 35567054
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 45,
                                "total_size": 3854100
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 1567,
                                "total_size": 1063419014
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 33,
                                "total_size": 3742164
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 12,
                                "total_size": 89155
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 490,
                                "total_size": 35888442
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 5498,
                                "total_size": 521390712
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 616,
                                "total_size": 43192392
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 115,
                                "total_size": 2054279
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 464,
                                "total_size": 14862314
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 550,
                                "total_size": 144149638
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 11,
                                "total_size": 629035
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 16,
                                "total_size": 640492
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 2669,
                                "total_size": 236862071
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 592,
                                "total_size": 143030161
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 41,
                                "total_size": 672310
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 861,
                                "total_size": 643612540
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 288,
                                "total_size": 190069707
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 401,
                                "total_size": 2451261786
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 30,
                                "total_size": 172367
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 790,
                                "total_size": 209010111
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 202,
                                "total_size": 122152715
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 516,
                                "total_size": 179155395
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 31,
                                "total_size": 9672751
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 3916,
                                "total_size": 245565907
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 4420,
                                "total_size": 350299656
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 925,
                                "total_size": 93982989
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 93,
                                "total_size": 3067135
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 942,
                                "total_size": 591527260
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 1044,
                                "total_size": 995495763
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 914,
                                "total_size": 235601761
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 676,
                                "total_size": 161246400
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 1022,
                                "total_size": 108154376
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 804,
                                "total_size": 238909157
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 336,
                                "total_size": 37921434
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 506,
                                "total_size": 149285502
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 6108,
                                "total_size": 623554565
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 500,
                                "total_size": 14395545
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 150,
                                "total_size": 108035107
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 490,
                                "total_size": 203244102
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 653,
                                "total_size": 214545054
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 507,
                                "total_size": 398454113
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 105,
                                "total_size": 189564257
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 8,
                                "total_size": 292959777
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 61,
                                "total_size": 37031476
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 551,
                                "total_size": 186385180
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 38,
                                "total_size": 978888
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 847,
                                "total_size": 2119144774
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 41,
                                "total_size": 507971826
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 8408,
                                "total_size": 779435490
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 264,
                                "total_size": 65109831
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 37,
                                "total_size": 4302609
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 15,
                                "total_size": 258756
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 414,
                                "total_size": 6301548
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 641,
                                "total_size": 87718707
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 707,
                                "total_size": 34122484
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 197,
                                "total_size": 8575427
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 527,
                                "total_size": 59371658
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 693,
                                "total_size": 62415522
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 45,
                                "total_size": 3524484
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 97,
                                "total_size": 1746294
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 1976,
                                "total_size": 1338317894
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 1105,
                                "total_size": 196842978
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 919,
                                "total_size": 761303027
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 663,
                                "total_size": 434679135
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 512,
                                "total_size": 206166030
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 6,
                                "total_size": 8817413
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 7935,
                                "total_size": 720805294
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 322,
                                "total_size": 105467552
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 320,
                                "total_size": 181741137
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 662,
                                "total_size": 266769045
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 1593,
                                "total_size": 645568654
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 304,
                                "total_size": 146166164
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 792,
                                "total_size": 81142379
                            },
                            {
                                "rc_agent": "c18624e7-e7f5-48c1-9352-4965a0f8c59e",
                                "total_count": 696,
                                "total_size": 61188944
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 663,
                                "total_size": 53729204
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 49,
                                "total_size": 720060860
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 135,
                                "total_size": 13503723
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 4181,
                                "total_size": 276107349
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 1433,
                                "total_size": 240714247
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 60,
                                "total_size": 18607932
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 412,
                                "total_size": 134911002
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 454,
                                "total_size": 229966505
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 530,
                                "total_size": 147813258
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 8,
                                "total_size": 2074016
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 396,
                                "total_size": 10866455
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 1230,
                                "total_size": 883755774
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 33,
                                "total_size": 26191948
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 1785,
                                "total_size": 238457653
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 1166,
                                "total_size": 247948677
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 526,
                                "total_size": 140138320
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 1288,
                                "total_size": 1007312918
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 628,
                                "total_size": 41911982
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 40,
                                "total_size": 13624543
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 127,
                                "total_size": 32560597
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 56,
                                "total_size": 1170780
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 704,
                                "total_size": 210692015
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 514,
                                "total_size": 50124342
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 125,
                                "total_size": 33997028
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 11417,
                                "total_size": 1080784007
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 529,
                                "total_size": 211585493
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 661,
                                "total_size": 599762752
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 53,
                                "total_size": 2419897
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 443,
                                "total_size": 75185440
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 1930,
                                "total_size": 658531551
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 16,
                                "total_size": 819751
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 62,
                                "total_size": 37193294
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 523,
                                "total_size": 26732026
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 753,
                                "total_size": 293407286
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 1378,
                                "total_size": 888461349
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 446,
                                "total_size": 114109851
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 860,
                                "total_size": 88915665
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 13045,
                                "total_size": 1172436679
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 2357,
                                "total_size": 629947290
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 155,
                                "total_size": 6249833
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 4464,
                                "total_size": 412342885
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 279,
                                "total_size": 10376299
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 634,
                                "total_size": 58114197
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 504,
                                "total_size": 410967628
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 7,
                                "total_size": 796672
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 5329,
                                "total_size": 451337851
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 6294,
                                "total_size": 571907199
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 1068,
                                "total_size": 47140402
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 3258,
                                "total_size": 254884139
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 354,
                                "total_size": 93305091
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 165,
                                "total_size": 102168351
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 283,
                                "total_size": 176536847
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 392,
                                "total_size": 1336709808
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 5504,
                                "total_size": 490254974
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 61,
                                "total_size": 136064903
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 869,
                                "total_size": 2678930941
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 167,
                                "total_size": 130341761
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 534,
                                "total_size": 778463898
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 534,
                                "total_size": 517049038
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 55,
                                "total_size": 660698
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 151,
                                "total_size": 146074969
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 543,
                                "total_size": 123576682
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 83,
                                "total_size": 4750844
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 4457,
                                "total_size": 300691201
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 8078,
                                "total_size": 710055595
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 34,
                                "total_size": 262800686
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 1057,
                                "total_size": 114095501
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 29,
                                "total_size": 1192391
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 674,
                                "total_size": 467540833
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 56,
                                "total_size": 22507140
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 845,
                                "total_size": 354169305
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 17,
                                "total_size": 665973
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 437,
                                "total_size": 227549829
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 2118,
                                "total_size": 2789083621
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 5053,
                                "total_size": 476230495
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 917,
                                "total_size": 1474523910
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 29,
                                "total_size": 314700
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 499,
                                "total_size": 19467011
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 367,
                                "total_size": 99472889
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 572,
                                "total_size": 201323334
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 588,
                                "total_size": 2321036698
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 470,
                                "total_size": 107392793
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 318,
                                "total_size": 35955924
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 1095,
                                "total_size": 322249513
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 175,
                                "total_size": 11083377
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 336,
                                "total_size": 51757752
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 2144,
                                "total_size": 1111598091
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 84,
                                "total_size": 216137948
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 912,
                                "total_size": 210885675
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 280,
                                "total_size": 120690068
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 44,
                                "total_size": 4067384
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 504,
                                "total_size": 40170271
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 51,
                                "total_size": 2214714
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 666,
                                "total_size": 604534625
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 515,
                                "total_size": 63732330
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 851,
                                "total_size": 87107686
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 6258,
                                "total_size": 564942218
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 138,
                                "total_size": 2698826
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 76,
                                "total_size": 755102
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 831,
                                "total_size": 229609341
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 627,
                                "total_size": 242100020
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 535,
                                "total_size": 128146956
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 15,
                                "total_size": 1505612
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 4,
                                "total_size": 1295484
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 5,
                                "total_size": 164933169
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 547,
                                "total_size": 187571608
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 507,
                                "total_size": 14674821
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 654,
                                "total_size": 48961492
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 898,
                                "total_size": 2794094176
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 2022,
                                "total_size": 1322155534
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 381,
                                "total_size": 24733756
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 5054,
                                "total_size": 428784868
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 628,
                                "total_size": 54129770
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 501,
                                "total_size": 104743864
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 151,
                                "total_size": 34367685
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 40,
                                "total_size": 30587440
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 5,
                                "total_size": 4160745
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 901,
                                "total_size": 89769664
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 1016,
                                "total_size": 719087778
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 534,
                                "total_size": 39923321
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 448,
                                "total_size": 112813421
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 467,
                                "total_size": 2865374576
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 586,
                                "total_size": 22912371
                            }
                        ]
                    },
                    "time": "2025-02-20T07:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "n5ViJZUBV5X_QYDfAJwV",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 638535,
                        "total_size": 185385535023,
                        "group_by_agent": [
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 239,
                                "total_size": 224452599
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 638,
                                "total_size": 509399564
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 751,
                                "total_size": 756556269
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 850,
                                "total_size": 202254433
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 1127,
                                "total_size": 336015433
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 12441,
                                "total_size": 1275649813
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 69,
                                "total_size": 12367878
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 11,
                                "total_size": 556055
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 44,
                                "total_size": 3235308
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 3874,
                                "total_size": 344542824
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 1681,
                                "total_size": 511729589
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 3465,
                                "total_size": 1933412096
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 1080,
                                "total_size": 975660234
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 1284,
                                "total_size": 89951255
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 309,
                                "total_size": 12489643
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 736,
                                "total_size": 284807236
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 92,
                                "total_size": 52898128
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 27,
                                "total_size": 3451427
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 4,
                                "total_size": 22284
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 986,
                                "total_size": 213561124
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 976,
                                "total_size": 1582218224
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 69,
                                "total_size": 756664
                            },
                            {
                                "rc_agent": "e4fdcdc5-31e2-4f25-89fd-79532caa8f26",
                                "total_count": 64,
                                "total_size": 42547953
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 1238,
                                "total_size": 472021574
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 658,
                                "total_size": 77739562
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 15,
                                "total_size": 596973
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 568,
                                "total_size": 18166424
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 2456,
                                "total_size": 250687518
                            },
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 1726,
                                "total_size": 433065105
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 2449,
                                "total_size": 527924944
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 841,
                                "total_size": 192968671
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 488,
                                "total_size": 302390907
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 46,
                                "total_size": 12559336
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 12107,
                                "total_size": 1100171520
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 85,
                                "total_size": 26654960
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 732,
                                "total_size": 192194608
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 334,
                                "total_size": 143207137
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 1497,
                                "total_size": 1038704674
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 148,
                                "total_size": 1491637
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 69,
                                "total_size": 54075887
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 1052,
                                "total_size": 703748082
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 721,
                                "total_size": 4563724288
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 970,
                                "total_size": 388930966
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 769,
                                "total_size": 56765518
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 11914,
                                "total_size": 1210794080
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 711,
                                "total_size": 113802971
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 188,
                                "total_size": 146614655
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 375,
                                "total_size": 271713864
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 969,
                                "total_size": 37753752
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 31,
                                "total_size": 46639782
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 15867,
                                "total_size": 1491009515
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 719,
                                "total_size": 151168858
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 826,
                                "total_size": 333347357
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 1315,
                                "total_size": 108867170
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 1839,
                                "total_size": 181452512
                            },
                            {
                                "rc_agent": "456d9bf9-a53a-4786-9c2c-24cdd5fd6771",
                                "total_count": 2704,
                                "total_size": 1201213107
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 302,
                                "total_size": 57072544
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 330,
                                "total_size": 23696984
                            },
                            {
                                "rc_agent": "59f0c72a-68e9-4cd8-a0d5-987ffe7fc0e1",
                                "total_count": 14,
                                "total_size": 4519730
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 715,
                                "total_size": 721424927
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 3121,
                                "total_size": 2218366061
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 796,
                                "total_size": 13109463
                            },
                            {
                                "rc_agent": "cde365e1-868f-493d-9adf-7ce505eec802",
                                "total_count": 1,
                                "total_size": 72654
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 1060,
                                "total_size": 432239868
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 2706,
                                "total_size": 225811139
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 1386,
                                "total_size": 474578367
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 175,
                                "total_size": 57125302
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 126,
                                "total_size": 341274831
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 86,
                                "total_size": 2163009
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 141,
                                "total_size": 43159357
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 827,
                                "total_size": 431707155
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 58,
                                "total_size": 35157753
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 993,
                                "total_size": 220270870
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 116,
                                "total_size": 2420815
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 28,
                                "total_size": 484147
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 119,
                                "total_size": 35966009
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 150,
                                "total_size": 4952595
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 1222,
                                "total_size": 495237311
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 312,
                                "total_size": 74151667
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 8824,
                                "total_size": 582182368
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 1138,
                                "total_size": 89228486
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 105,
                                "total_size": 17845262
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 3867,
                                "total_size": 7253034361
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 10945,
                                "total_size": 945573871
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 18501,
                                "total_size": 1804069244
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 681,
                                "total_size": 225950156
                            },
                            {
                                "rc_agent": "d5e8e4a3-2d59-4e1d-9e37-b0917e7dd330",
                                "total_count": 668,
                                "total_size": 561319865
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 646,
                                "total_size": 226053933
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 1047,
                                "total_size": 436486011
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 234,
                                "total_size": 63664860
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 9984,
                                "total_size": 886629727
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 44,
                                "total_size": 4553004
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 3611,
                                "total_size": 1756473297
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 2,
                                "total_size": 2311
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 8044,
                                "total_size": 543851112
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 11530,
                                "total_size": 979376999
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 26,
                                "total_size": 2740540
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 714,
                                "total_size": 678516490
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 1713,
                                "total_size": 443781033
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 1721,
                                "total_size": 4332810474
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 76,
                                "total_size": 3180967
                            },
                            {
                                "rc_agent": "c5d364ec-ebe9-46d2-b54b-09fbdad41ac8",
                                "total_count": 750,
                                "total_size": 222659954
                            },
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 3196,
                                "total_size": 308799935
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 1867,
                                "total_size": 160782407
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 98,
                                "total_size": 938800888
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 487,
                                "total_size": 18285482
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 20,
                                "total_size": 1033076
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 11970,
                                "total_size": 748987152
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 21456,
                                "total_size": 2080562289
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 652,
                                "total_size": 148263530
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 1734,
                                "total_size": 5251131971
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 101,
                                "total_size": 22709404
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 2494,
                                "total_size": 7493022396
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 72,
                                "total_size": 17092166
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 956,
                                "total_size": 333592429
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 1295,
                                "total_size": 97335731
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 345,
                                "total_size": 215124962
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 815,
                                "total_size": 299255745
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 754,
                                "total_size": 716235377
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 4,
                                "total_size": 121178
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 7328,
                                "total_size": 650149414
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 24238,
                                "total_size": 2158871417
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 1008,
                                "total_size": 220763888
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 390,
                                "total_size": 257010258
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 74,
                                "total_size": 5176044
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 334,
                                "total_size": 105449861
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 1292,
                                "total_size": 96502821
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 1103,
                                "total_size": 181218371
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 75,
                                "total_size": 57531501
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 1414,
                                "total_size": 4133234876
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 3919,
                                "total_size": 317974083
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 1317,
                                "total_size": 326395176
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 351,
                                "total_size": 15405127
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 623,
                                "total_size": 65497502
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 1336,
                                "total_size": 1073524462
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 3111,
                                "total_size": 1514679153
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 3611,
                                "total_size": 2509016014
                            },
                            {
                                "rc_agent": "0143a74d-9e52-4478-aba0-f0853a0d9f02",
                                "total_count": 1073,
                                "total_size": 254055813
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 926,
                                "total_size": 74916807
                            },
                            {
                                "rc_agent": "7fe8e59a-b5c1-4e36-a9f4-b526023335fc",
                                "total_count": 18277,
                                "total_size": 1942091553
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 154,
                                "total_size": 86919827
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 95,
                                "total_size": 44179371
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 512,
                                "total_size": 192925815
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 13161,
                                "total_size": 1219009366
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 888,
                                "total_size": 71809480
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 770,
                                "total_size": 197148326
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 1324,
                                "total_size": 60522113
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 61,
                                "total_size": 7193408
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 840,
                                "total_size": 114659806
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 11064,
                                "total_size": 1006254794
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 54,
                                "total_size": 24616836
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 1434,
                                "total_size": 131245222
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 691,
                                "total_size": 64593988
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 451,
                                "total_size": 277124822
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 168,
                                "total_size": 56819039
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 799,
                                "total_size": 270029677
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 333,
                                "total_size": 89965899
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 14581,
                                "total_size": 1237844068
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 2589,
                                "total_size": 1180805385
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 2878,
                                "total_size": 5951219052
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 83,
                                "total_size": 1263427512
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 265,
                                "total_size": 146627568
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 177,
                                "total_size": 7913841
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 5683,
                                "total_size": 522202689
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 1453,
                                "total_size": 432738994
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 1317,
                                "total_size": 92225017
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 48,
                                "total_size": 78970889
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 995,
                                "total_size": 6201429526
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 753,
                                "total_size": 99061638
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 757,
                                "total_size": 2581350318
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 8961,
                                "total_size": 850620373
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 945,
                                "total_size": 59024842
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 724,
                                "total_size": 284429851
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 1208,
                                "total_size": 86645801
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 734,
                                "total_size": 52391483
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 627,
                                "total_size": 166973035
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 946,
                                "total_size": 27265474
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 19,
                                "total_size": 598469875
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 664,
                                "total_size": 158705058
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 763,
                                "total_size": 230594880
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 19,
                                "total_size": 1137542
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 13099,
                                "total_size": 1038563428
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 1222,
                                "total_size": 92188940
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 82,
                                "total_size": 13619314
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 2,
                                "total_size": 1711421
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 13816,
                                "total_size": 1296742923
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 652,
                                "total_size": 506675624
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 11,
                                "total_size": 81785
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 1209,
                                "total_size": 355969879
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 2172,
                                "total_size": 1415850274
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 1575,
                                "total_size": 1473050796
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 232,
                                "total_size": 4544522
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 1117,
                                "total_size": 786481533
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 10857,
                                "total_size": 1060731356
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 1651,
                                "total_size": 432302200
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 718,
                                "total_size": 88610453
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 6,
                                "total_size": 18429533
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 42,
                                "total_size": 717792
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 1407,
                                "total_size": 126637309
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 206,
                                "total_size": 3740338
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 31,
                                "total_size": 16615353
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 682,
                                "total_size": 76979722
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 89,
                                "total_size": 148999137
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 11322,
                                "total_size": 1142195767
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 92,
                                "total_size": 67823755
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 932,
                                "total_size": 36903109
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 815,
                                "total_size": 322406129
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 16,
                                "total_size": 1840208
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 70,
                                "total_size": 7973379
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 3495,
                                "total_size": 1647284080
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 877,
                                "total_size": 25258011
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 145,
                                "total_size": 13030991
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 870,
                                "total_size": 215989881
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 721,
                                "total_size": 472739657
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 1397,
                                "total_size": 103563276
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 2725,
                                "total_size": 1544034018
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 2545,
                                "total_size": 260689463
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 1174,
                                "total_size": 85854791
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 1466,
                                "total_size": 1123463297
                            },
                            {
                                "rc_agent": "e5bfa1ac-6282-4695-b99f-b948c9244adc",
                                "total_count": 9950,
                                "total_size": 888263920
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 18323,
                                "total_size": 1683962000
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 63,
                                "total_size": 150348561
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 1748,
                                "total_size": 5374518045
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 807,
                                "total_size": 42428203
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 823,
                                "total_size": 412518837
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 766,
                                "total_size": 78803606
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 57,
                                "total_size": 1475293
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 487,
                                "total_size": 160839549
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 31,
                                "total_size": 575034
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 29,
                                "total_size": 17844927
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 87,
                                "total_size": 1007782
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 902,
                                "total_size": 16521264
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 233,
                                "total_size": 22927588
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 21002,
                                "total_size": 1953405458
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 771,
                                "total_size": 184918692
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 375,
                                "total_size": 178766803
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 12981,
                                "total_size": 1232063108
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 909,
                                "total_size": 81624661
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 2302,
                                "total_size": 1687100266
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 397,
                                "total_size": 307702793
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 24,
                                "total_size": 567908077
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 90,
                                "total_size": 8835515
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 194,
                                "total_size": 118494394
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 530,
                                "total_size": 293482225
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 3646,
                                "total_size": 2124015409
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 787,
                                "total_size": 39558890
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 751,
                                "total_size": 102403089
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 71,
                                "total_size": 3254633
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 1149,
                                "total_size": 4619001564
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 55,
                                "total_size": 1401251
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 1781,
                                "total_size": 451229116
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 258,
                                "total_size": 147732182
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 1225,
                                "total_size": 310832425
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 2864,
                                "total_size": 378115377
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 747,
                                "total_size": 33898476
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 713,
                                "total_size": 1040284142
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 70,
                                "total_size": 23886502
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 207,
                                "total_size": 162904184
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 36,
                                "total_size": 1475135
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 1302,
                                "total_size": 119402739
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 948,
                                "total_size": 374517190
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 648,
                                "total_size": 4048232963
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 72,
                                "total_size": 903181407
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 4156,
                                "total_size": 2799759497
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 2016,
                                "total_size": 956585829
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 1225,
                                "total_size": 94335674
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 1353,
                                "total_size": 3812218390
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 1477,
                                "total_size": 373156484
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 1224,
                                "total_size": 642183602
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 13677,
                                "total_size": 1203290871
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 8116,
                                "total_size": 635673143
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 1573,
                                "total_size": 384583908
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 68,
                                "total_size": 7056071
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 1113,
                                "total_size": 329357046
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 769,
                                "total_size": 87580227
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 2661,
                                "total_size": 264045440
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 678,
                                "total_size": 18751353
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 792,
                                "total_size": 219691440
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 27,
                                "total_size": 5183549
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 799,
                                "total_size": 66759298
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 37,
                                "total_size": 543706527
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 53,
                                "total_size": 1411843
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 18,
                                "total_size": 1017245
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 401,
                                "total_size": 227515318
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 4,
                                "total_size": 3220115
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 755,
                                "total_size": 204767931
                            },
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 3433,
                                "total_size": 276889026
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 9855,
                                "total_size": 984383383
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 1072,
                                "total_size": 440583507
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 106,
                                "total_size": 28021869
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 3716,
                                "total_size": 5028983242
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 2940,
                                "total_size": 1487363311
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 2525,
                                "total_size": 1830243394
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 1117,
                                "total_size": 81051749
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 89,
                                "total_size": 168302078
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 58,
                                "total_size": 13265479
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 16,
                                "total_size": 1524442
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 748,
                                "total_size": 197645488
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 2114,
                                "total_size": 1350126946
                            }
                        ]
                    },
                    "time": "2025-02-20T03:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "sJViJZUBV5X_QYDfAJwV",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 889099,
                        "total_size": 201886393263,
                        "group_by_agent": [
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 20612,
                                "total_size": 1869303480
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 752,
                                "total_size": 418750339
                            },
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 3747,
                                "total_size": 933312590
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 1032,
                                "total_size": 805856827
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 536,
                                "total_size": 409506421
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 536,
                                "total_size": 428910513
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 72,
                                "total_size": 22592743
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 6048,
                                "total_size": 546017384
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 624,
                                "total_size": 165738927
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 52,
                                "total_size": 3177830
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 611,
                                "total_size": 231552083
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 1033,
                                "total_size": 222730204
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 81,
                                "total_size": 25065370
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 5139,
                                "total_size": 461236745
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 868,
                                "total_size": 65025711
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 6825,
                                "total_size": 4747309137
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 1432,
                                "total_size": 128571620
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 1899,
                                "total_size": 1102591237
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 48,
                                "total_size": 620924
                            },
                            {
                                "rc_agent": "7e515c16-ca38-4323-afd6-53d7b79c74f3",
                                "total_count": 112,
                                "total_size": 11811231
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 5022,
                                "total_size": 2384694364
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 2051,
                                "total_size": 594714911
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 1114,
                                "total_size": 282188691
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 14454,
                                "total_size": 1429821187
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 114,
                                "total_size": 1972866
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 432,
                                "total_size": 2698710076
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 548,
                                "total_size": 273586452
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 40,
                                "total_size": 2309078
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 249,
                                "total_size": 151282501
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 224,
                                "total_size": 22672399
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 262,
                                "total_size": 4723400
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 160,
                                "total_size": 261942765
                            },
                            {
                                "rc_agent": "822d1955-483d-4143-a904-e722cf12a0d4",
                                "total_count": 112,
                                "total_size": 26706911
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 299,
                                "total_size": 177264551
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 21800,
                                "total_size": 1439442997
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 131,
                                "total_size": 307705691
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 648,
                                "total_size": 217647467
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 536,
                                "total_size": 121688944
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 696,
                                "total_size": 37558606
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 1990,
                                "total_size": 475530685
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 1560,
                                "total_size": 125062221
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 2964,
                                "total_size": 2133296403
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 2398,
                                "total_size": 178582374
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 96,
                                "total_size": 11426411
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 4991,
                                "total_size": 3352076708
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 19491,
                                "total_size": 1812997569
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 144,
                                "total_size": 936031018
                            },
                            {
                                "rc_agent": "75e33732-2a28-4529-bf91-9b17995e358f",
                                "total_count": 18304,
                                "total_size": 808206247
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 2148,
                                "total_size": 536935691
                            },
                            {
                                "rc_agent": "0143a74d-9e52-4478-aba0-f0853a0d9f02",
                                "total_count": 2004,
                                "total_size": 462169849
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 2464,
                                "total_size": 175075223
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 1040,
                                "total_size": 640449948
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 3034,
                                "total_size": 1235845049
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 14329,
                                "total_size": 1462894778
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 864,
                                "total_size": 299258718
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 648,
                                "total_size": 70659676
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 1658,
                                "total_size": 219718792
                            },
                            {
                                "rc_agent": "9f198d32-faae-441f-899f-9587166f5202",
                                "total_count": 1426,
                                "total_size": 313950033
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 22989,
                                "total_size": 2033857658
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 21457,
                                "total_size": 1915097808
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 6471,
                                "total_size": 2936257006
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 850,
                                "total_size": 23293253
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 210,
                                "total_size": 8458425
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 26,
                                "total_size": 439163
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 227,
                                "total_size": 22721210
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 688,
                                "total_size": 1715168516
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 2,
                                "total_size": 1536751
                            },
                            {
                                "rc_agent": "456d9bf9-a53a-4786-9c2c-24cdd5fd6771",
                                "total_count": 5021,
                                "total_size": 2247624643
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 18795,
                                "total_size": 1473166756
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 80,
                                "total_size": 62602902
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 464,
                                "total_size": 2900191321
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 68,
                                "total_size": 7765891
                            },
                            {
                                "rc_agent": "5bbbb159-b28d-4c13-9d0f-636d8484fce6",
                                "total_count": 136,
                                "total_size": 64495691
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 1324,
                                "total_size": 38231587
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 3584,
                                "total_size": 760717135
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 2080,
                                "total_size": 1940770742
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 381,
                                "total_size": 234800852
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 313,
                                "total_size": 305278422
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 476,
                                "total_size": 346479856
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 616,
                                "total_size": 3843143736
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 872,
                                "total_size": 295077149
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 1361,
                                "total_size": 43667512
                            },
                            {
                                "rc_agent": "6c3dd642-42f6-42b6-ac74-3142252b4e57",
                                "total_count": 6576,
                                "total_size": 655602397
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 48,
                                "total_size": 2502030
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 680,
                                "total_size": 10336430
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 632,
                                "total_size": 149658510
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 291,
                                "total_size": 182107097
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 32,
                                "total_size": 5494841
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 56,
                                "total_size": 2314049
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 1031,
                                "total_size": 688734522
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 1674,
                                "total_size": 386980595
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 1010,
                                "total_size": 392354712
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 408,
                                "total_size": 146136630
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 59,
                                "total_size": 14866220
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 1310,
                                "total_size": 690689274
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 19208,
                                "total_size": 1712070001
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 2000,
                                "total_size": 514728292
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 120,
                                "total_size": 1422277
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 4760,
                                "total_size": 3361247399
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 3464,
                                "total_size": 1643223774
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 13969,
                                "total_size": 1423726743
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 3864,
                                "total_size": 2791630707
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 5136,
                                "total_size": 2333520524
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 1806,
                                "total_size": 945778270
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 922,
                                "total_size": 1343121478
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 3942,
                                "total_size": 2133654783
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 66,
                                "total_size": 108707346
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 818,
                                "total_size": 2789358732
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 72,
                                "total_size": 876330030
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 840,
                                "total_size": 63188571
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 848,
                                "total_size": 77652498
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 1120,
                                "total_size": 82499016
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 315,
                                "total_size": 245329182
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 1479,
                                "total_size": 4193972418
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 1016,
                                "total_size": 29279070
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 912,
                                "total_size": 733471524
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 100,
                                "total_size": 575468337
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 29,
                                "total_size": 227498042
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 1346,
                                "total_size": 52293817
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 15826,
                                "total_size": 1487656510
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 954,
                                "total_size": 199824565
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 2416,
                                "total_size": 4874158619
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 248,
                                "total_size": 246178983
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 28,
                                "total_size": 3593623
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 47,
                                "total_size": 2025074
                            },
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 3382,
                                "total_size": 273571941
                            },
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 4795,
                                "total_size": 503611778
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 688,
                                "total_size": 48984000
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 1080,
                                "total_size": 740049841
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 1753,
                                "total_size": 7145520681
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 540,
                                "total_size": 71939236
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 418,
                                "total_size": 237855325
                            },
                            {
                                "rc_agent": "f723bfbb-36bd-46f7-be3e-a53c9b71b3ad",
                                "total_count": 2697,
                                "total_size": 176984026
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 1875,
                                "total_size": 474408644
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 20106,
                                "total_size": 1715126313
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 1871,
                                "total_size": 482041396
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 2178,
                                "total_size": 545007125
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 1091,
                                "total_size": 85869102
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 736,
                                "total_size": 175923884
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 6,
                                "total_size": 7086
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 12363,
                                "total_size": 1122233250
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 2653,
                                "total_size": 229012250
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 4512,
                                "total_size": 1546493241
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 819,
                                "total_size": 210367761
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 1282,
                                "total_size": 383841344
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 272,
                                "total_size": 73920661
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 8316,
                                "total_size": 518475554
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 1936,
                                "total_size": 91093324
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 128,
                                "total_size": 3477081
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 997,
                                "total_size": 297641730
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 152,
                                "total_size": 70804476
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 524,
                                "total_size": 308738698
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 4880,
                                "total_size": 2135268453
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 1308,
                                "total_size": 428598817
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 13234,
                                "total_size": 1246312349
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 1454,
                                "total_size": 92226520
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 572,
                                "total_size": 229010553
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 257,
                                "total_size": 69785482
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 14328,
                                "total_size": 1331544382
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 1481,
                                "total_size": 574975509
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 136,
                                "total_size": 345907192
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 18,
                                "total_size": 1005077
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 664,
                                "total_size": 49726186
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 24238,
                                "total_size": 2095152130
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 1376,
                                "total_size": 2477350007
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 26472,
                                "total_size": 2576808956
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 200,
                                "total_size": 6672182
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 158,
                                "total_size": 103984034
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 113,
                                "total_size": 1126549
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 544,
                                "total_size": 142148329
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 636,
                                "total_size": 81184480
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 16765,
                                "total_size": 1362699995
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 2648,
                                "total_size": 234215219
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 80,
                                "total_size": 2034107
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 688,
                                "total_size": 188638373
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 10340,
                                "total_size": 692516920
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 156,
                                "total_size": 62411021
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 1264,
                                "total_size": 98346448
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 1586,
                                "total_size": 685179878
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 738,
                                "total_size": 280022751
                            },
                            {
                                "rc_agent": "2dd609cf-1d99-4f2b-93f5-06037a040d67",
                                "total_count": 49,
                                "total_size": 24041648
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 1611,
                                "total_size": 482765975
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 88,
                                "total_size": 16931429
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 4545,
                                "total_size": 2545489361
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 3464,
                                "total_size": 4678028443
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 3630,
                                "total_size": 1098692847
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 1036,
                                "total_size": 214112949
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 680,
                                "total_size": 471963707
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 352,
                                "total_size": 48207534
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 19777,
                                "total_size": 1571754491
                            },
                            {
                                "rc_agent": "8e572aba-c152-4b25-9b98-696042bae8b7",
                                "total_count": 1421,
                                "total_size": 134234416
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 892,
                                "total_size": 497698056
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 588,
                                "total_size": 23095387
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 138,
                                "total_size": 5877003
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 732,
                                "total_size": 288637091
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 352,
                                "total_size": 40209022
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 44,
                                "total_size": 780926896
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 2200,
                                "total_size": 1490513726
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 108,
                                "total_size": 2588548
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 211,
                                "total_size": 4078653
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 656,
                                "total_size": 1992699087
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 11808,
                                "total_size": 1169641006
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 1560,
                                "total_size": 124952450
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 1170,
                                "total_size": 398452369
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 984,
                                "total_size": 799959323
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 464,
                                "total_size": 53790157
                            },
                            {
                                "rc_agent": "dcef71ee-d373-4ed5-81cf-780f80674f3d",
                                "total_count": 991,
                                "total_size": 731089532
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 96,
                                "total_size": 33047394
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 36,
                                "total_size": 481735
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 792,
                                "total_size": 187196295
                            },
                            {
                                "rc_agent": "e5bfa1ac-6282-4695-b99f-b948c9244adc",
                                "total_count": 17421,
                                "total_size": 1548362596
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 19222,
                                "total_size": 1770035488
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 24,
                                "total_size": 911177
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 463,
                                "total_size": 33517558
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 1560,
                                "total_size": 124368007
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 880,
                                "total_size": 120127688
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 572,
                                "total_size": 109659790
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 33793,
                                "total_size": 3098966388
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 1804,
                                "total_size": 5358826664
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 664,
                                "total_size": 61023549
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 872,
                                "total_size": 40040316
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 864,
                                "total_size": 202423699
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 2948,
                                "total_size": 294920943
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 984,
                                "total_size": 605851004
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 680,
                                "total_size": 2054440575
                            },
                            {
                                "rc_agent": "bace3ddf-e238-4660-9232-4d661fe426bf",
                                "total_count": 2008,
                                "total_size": 187340270
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 96,
                                "total_size": 58432672
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 1264,
                                "total_size": 95693676
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 1761,
                                "total_size": 387309125
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 952,
                                "total_size": 928536566
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 1367,
                                "total_size": 101129387
                            },
                            {
                                "rc_agent": "c9c2bb7c-83ec-4fbd-b573-b157e80bf82b",
                                "total_count": 48,
                                "total_size": 7877621
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 88,
                                "total_size": 9381784
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 936,
                                "total_size": 152319154
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 101,
                                "total_size": 65169332
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 246,
                                "total_size": 161102795
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 18144,
                                "total_size": 1762223904
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 1059,
                                "total_size": 2963372427
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 1862,
                                "total_size": 3444576740
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 84,
                                "total_size": 64278445
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 1893,
                                "total_size": 578344080
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 2200,
                                "total_size": 1414347206
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 18,
                                "total_size": 2045689
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 1407,
                                "total_size": 427683903
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 836,
                                "total_size": 58457064
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 2606,
                                "total_size": 232948921
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 772,
                                "total_size": 257063481
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 144,
                                "total_size": 353657487
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 1080,
                                "total_size": 953065520
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 688,
                                "total_size": 61990573
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 632,
                                "total_size": 32137689
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 944,
                                "total_size": 733758999
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 1450,
                                "total_size": 347016658
                            },
                            {
                                "rc_agent": "0c98c7b5-f0bf-4b03-a667-a754e5acefd3",
                                "total_count": 29856,
                                "total_size": 2623657659
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 26258,
                                "total_size": 2515723127
                            },
                            {
                                "rc_agent": "7fe8e59a-b5c1-4e36-a9f4-b526023335fc",
                                "total_count": 25546,
                                "total_size": 2714951159
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 508,
                                "total_size": 19121398
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 9925,
                                "total_size": 827903022
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 575,
                                "total_size": 235206866
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 891,
                                "total_size": 366069205
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 855,
                                "total_size": 827530110
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 2464,
                                "total_size": 206281452
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 5,
                                "total_size": 262762
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 6,
                                "total_size": 190394892
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 995,
                                "total_size": 517890350
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 2464,
                                "total_size": 209786511
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 756,
                                "total_size": 204893808
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 334,
                                "total_size": 143277293
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 88,
                                "total_size": 15707989
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 338,
                                "total_size": 105860026
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 62,
                                "total_size": 1590833
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 61,
                                "total_size": 346247
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 248,
                                "total_size": 238538628
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 62,
                                "total_size": 1710780
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 15449,
                                "total_size": 1462444520
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 630,
                                "total_size": 155969919
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 952,
                                "total_size": 959720896
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 340,
                                "total_size": 32158361
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 464,
                                "total_size": 55383504
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 2172,
                                "total_size": 362566007
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 1328,
                                "total_size": 114118875
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 1482,
                                "total_size": 491897539
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 48,
                                "total_size": 72054437
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 48,
                                "total_size": 1916238
                            }
                        ]
                    },
                    "time": "2025-02-20T00:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "sZViJZUBV5X_QYDfAJwV",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-19T16:00:00.000000Z",
                    "indices": {
                        "total_count": 737821,
                        "total_size": 219740252482,
                        "group_by_agent": [
                            {
                                "rc_agent": "4c88f2ee-b61c-4053-bc85-84c5bb63e708",
                                "total_count": 1675,
                                "total_size": 420732986
                            },
                            {
                                "rc_agent": "02049b25-9286-429f-b86b-ee5b13973519",
                                "total_count": 1516,
                                "total_size": 456956237
                            },
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 1075,
                                "total_size": 79329621
                            },
                            {
                                "rc_agent": "cfc1eeef-c752-40d0-95bd-35c868f3e664",
                                "total_count": 1125,
                                "total_size": 331906696
                            },
                            {
                                "rc_agent": "a6dcaa15-0a8a-4b83-bf6a-e7f9f0cc05f5",
                                "total_count": 1169,
                                "total_size": 85653206
                            },
                            {
                                "rc_agent": "b3b9e139-c831-491f-8702-d6b9b92177f0",
                                "total_count": 1465,
                                "total_size": 110926330
                            },
                            {
                                "rc_agent": "adf9729e-4459-4892-bde9-72a413b818f8",
                                "total_count": 1211,
                                "total_size": 1218215825
                            },
                            {
                                "rc_agent": "ca6fd273-1408-4031-b609-8609dc75f1b9",
                                "total_count": 432,
                                "total_size": 18980700
                            },
                            {
                                "rc_agent": "14fb65cd-1336-435f-9db2-bc06d180ea5e",
                                "total_count": 127,
                                "total_size": 3224594
                            },
                            {
                                "rc_agent": "b225702d-5151-491d-ae0b-d6ffd8ed19e9",
                                "total_count": 376,
                                "total_size": 45603278
                            },
                            {
                                "rc_agent": "7dc27998-ea41-4c5c-9fde-d643589ccfc2",
                                "total_count": 4408,
                                "total_size": 581108940
                            },
                            {
                                "rc_agent": "73c7b282-d054-4fed-a200-7871f355d315",
                                "total_count": 62,
                                "total_size": 1575857
                            },
                            {
                                "rc_agent": "fd47524d-8208-4568-b55c-16790bd0da1e",
                                "total_count": 1145,
                                "total_size": 36302576
                            },
                            {
                                "rc_agent": "a0f1963b-6878-449f-a614-68cf08a88cbf",
                                "total_count": 984,
                                "total_size": 264160171
                            },
                            {
                                "rc_agent": "9e085a36-b19f-4834-9cf5-14907272b870",
                                "total_count": 1136,
                                "total_size": 397698561
                            },
                            {
                                "rc_agent": "311c5eae-5301-4ddc-bd7b-1d850cf268c9",
                                "total_count": 1192,
                                "total_size": 282628555
                            },
                            {
                                "rc_agent": "bc4592c7-e176-4334-853c-18466e6b488d",
                                "total_count": 182,
                                "total_size": 8169113
                            },
                            {
                                "rc_agent": "cf44c7d6-b5f6-4367-85d4-087bec0d2b77",
                                "total_count": 84,
                                "total_size": 3554320
                            },
                            {
                                "rc_agent": "e1998571-a692-45b3-b2a9-87454185ee93",
                                "total_count": 45,
                                "total_size": 67636174
                            },
                            {
                                "rc_agent": "c980c78c-e02c-4110-a00c-bf5580647d81",
                                "total_count": 1000,
                                "total_size": 126065837
                            },
                            {
                                "rc_agent": "7727011a-372b-498f-b994-cda9c0eb9b25",
                                "total_count": 40,
                                "total_size": 7816056
                            },
                            {
                                "rc_agent": "415a4187-cd24-4868-92f5-480e134d0a62",
                                "total_count": 373,
                                "total_size": 101360687
                            },
                            {
                                "rc_agent": "32dbf3c7-c8be-4aa5-855e-e9ab11016170",
                                "total_count": 1108,
                                "total_size": 461824903
                            },
                            {
                                "rc_agent": "45d4941a-492c-420a-a4d0-8f088d06e83f",
                                "total_count": 21242,
                                "total_size": 2120553239
                            },
                            {
                                "rc_agent": "5599e67c-c8ea-42e9-9604-ccc0fabf0863",
                                "total_count": 68,
                                "total_size": 7906645
                            },
                            {
                                "rc_agent": "4ac33af1-213b-4429-bd36-bc1ab6429269",
                                "total_count": 155,
                                "total_size": 58764990
                            },
                            {
                                "rc_agent": "db655884-1bb3-468e-a1c5-70e6f651ec3b",
                                "total_count": 1112,
                                "total_size": 395285558
                            },
                            {
                                "rc_agent": "e5b32761-3b1e-49bf-82b8-afcabcb821a2",
                                "total_count": 1080,
                                "total_size": 839171492
                            },
                            {
                                "rc_agent": "f5f69672-7163-4c5b-8822-527ea4476b96",
                                "total_count": 225,
                                "total_size": 60871672
                            },
                            {
                                "rc_agent": "23da683e-3dd7-4778-b0bf-35a051d1f2a0",
                                "total_count": 15,
                                "total_size": 4089138
                            },
                            {
                                "rc_agent": "7fa375d6-9cb1-4df5-a34c-8155aeb0b2ab",
                                "total_count": 288,
                                "total_size": 690447594
                            },
                            {
                                "rc_agent": "2950e9df-6c76-48ec-ad35-59df4059833e",
                                "total_count": 1567,
                                "total_size": 648471724
                            },
                            {
                                "rc_agent": "b84398ad-67c8-4993-ba9f-d084f701ee3f",
                                "total_count": 732,
                                "total_size": 11144861
                            },
                            {
                                "rc_agent": "b24c3471-614b-4e78-9d8f-f869417af23b",
                                "total_count": 7040,
                                "total_size": 645633021
                            },
                            {
                                "rc_agent": "e5bfa1ac-6282-4695-b99f-b948c9244adc",
                                "total_count": 15021,
                                "total_size": 1340686254
                            },
                            {
                                "rc_agent": "3610ea5e-1ff7-4417-8b48-134caa2c6845",
                                "total_count": 1211,
                                "total_size": 86814294
                            },
                            {
                                "rc_agent": "b47fc995-1d35-4995-9897-549bdf3ae514",
                                "total_count": 1448,
                                "total_size": 132916199
                            },
                            {
                                "rc_agent": "1fce3308-a3ec-479e-9c44-f15496e6c003",
                                "total_count": 165,
                                "total_size": 12229005
                            },
                            {
                                "rc_agent": "aee0dd64-6f45-4a68-ba1e-69ea4dec21e0",
                                "total_count": 776,
                                "total_size": 92703166
                            },
                            {
                                "rc_agent": "6e4c3f5a-9afb-4cd4-89ec-0d2bced597bb",
                                "total_count": 430,
                                "total_size": 311713661
                            },
                            {
                                "rc_agent": "306caf1d-aa66-4e9e-a20f-96a8db25ec2f",
                                "total_count": 62,
                                "total_size": 355960
                            },
                            {
                                "rc_agent": "5d373bc8-229e-4403-a406-0ed1ead5c319",
                                "total_count": 42,
                                "total_size": 141904
                            },
                            {
                                "rc_agent": "cde365e1-868f-493d-9adf-7ce505eec802",
                                "total_count": 1,
                                "total_size": 863401
                            },
                            {
                                "rc_agent": "8cfbfe47-4d59-43e3-9815-61e00937aa1d",
                                "total_count": 2879,
                                "total_size": 1408565438
                            },
                            {
                                "rc_agent": "9e0a5690-8c73-4008-b39b-94a85a55fa58",
                                "total_count": 1319,
                                "total_size": 333497257
                            },
                            {
                                "rc_agent": "704afd79-88be-4a41-93f4-e2e729c248bf",
                                "total_count": 1212,
                                "total_size": 1222074648
                            },
                            {
                                "rc_agent": "f24a5fc8-8103-4279-a763-d02c1aa03e6c",
                                "total_count": 595,
                                "total_size": 376503134
                            },
                            {
                                "rc_agent": "cbbb4199-71a4-4689-90b2-9bd21d6cefb9",
                                "total_count": 1398,
                                "total_size": 40381702
                            },
                            {
                                "rc_agent": "6988c630-f21f-49fc-a513-204e1cacacbd",
                                "total_count": 10,
                                "total_size": 521840
                            },
                            {
                                "rc_agent": "75002c57-e4fe-4468-b03a-121026308b95",
                                "total_count": 1,
                                "total_size": 1551
                            },
                            {
                                "rc_agent": "05ebdbe2-72c7-4ed3-af4b-ac991666a438",
                                "total_count": 630,
                                "total_size": 359188204
                            },
                            {
                                "rc_agent": "6fb24615-bcea-4333-a0fe-3e0bba70f9ec",
                                "total_count": 11173,
                                "total_size": 967244099
                            },
                            {
                                "rc_agent": "fcf5d92e-a056-4a95-9422-013edbac7eb9",
                                "total_count": 4975,
                                "total_size": 3610178983
                            },
                            {
                                "rc_agent": "eb195d39-20a2-4861-a995-214fc873f3a3",
                                "total_count": 1357,
                                "total_size": 3814733383
                            },
                            {
                                "rc_agent": "9f68e7ee-50e3-4df7-a365-dc92cd83e8d9",
                                "total_count": 1486,
                                "total_size": 106618940
                            },
                            {
                                "rc_agent": "1a9d9cd5-a7a3-495a-909c-5d0c1d9e204f",
                                "total_count": 16273,
                                "total_size": 1662150794
                            },
                            {
                                "rc_agent": "d9b9d8b7-2eec-47bc-a6f6-0fd276f73dea",
                                "total_count": 20,
                                "total_size": 3008831
                            },
                            {
                                "rc_agent": "7bd3c05c-b6e6-4b2f-822c-796ac225d46d",
                                "total_count": 39,
                                "total_size": 4139010
                            },
                            {
                                "rc_agent": "486b8ed5-2035-464b-a553-3f0d6a666a80",
                                "total_count": 1733,
                                "total_size": 906274348
                            },
                            {
                                "rc_agent": "95e00a7b-4881-4e99-b12a-eefc57db618a",
                                "total_count": 30,
                                "total_size": 1196782
                            },
                            {
                                "rc_agent": "b688dcf7-922a-4e92-be66-aa5f7fba0d36",
                                "total_count": 1248,
                                "total_size": 283576888
                            },
                            {
                                "rc_agent": "71ad7688-64f0-41c9-b767-e92974094907",
                                "total_count": 15,
                                "total_size": 11790580
                            },
                            {
                                "rc_agent": "62a076a4-73a8-4a1f-a347-c9f7b87d23e3",
                                "total_count": 408,
                                "total_size": 272629961
                            },
                            {
                                "rc_agent": "da3c9262-caa8-4db5-89fc-a63f2fe7c242",
                                "total_count": 20906,
                                "total_size": 1837621873
                            },
                            {
                                "rc_agent": "b51c799e-a18b-4488-86e4-a52bf4822014",
                                "total_count": 1716,
                                "total_size": 527074787
                            },
                            {
                                "rc_agent": "84c86c00-b47f-4e83-9adf-f978f5485048",
                                "total_count": 1268,
                                "total_size": 101316728
                            },
                            {
                                "rc_agent": "e81eb1ea-1c7c-4a03-8fc4-227d7862d4db",
                                "total_count": 1018,
                                "total_size": 253856756
                            },
                            {
                                "rc_agent": "0425a5b9-8b91-4678-9a8c-3c6b4f39f6c5",
                                "total_count": 677,
                                "total_size": 270573811
                            },
                            {
                                "rc_agent": "ebb95aa0-f906-45ff-8683-b5af18a0ef9e",
                                "total_count": 78,
                                "total_size": 1986390
                            },
                            {
                                "rc_agent": "33fd57fb-df7c-4290-acc4-b2af66c483c5",
                                "total_count": 702,
                                "total_size": 390205375
                            },
                            {
                                "rc_agent": "db8ca861-137b-434e-bfb7-cb1c665d023f",
                                "total_count": 1376,
                                "total_size": 126432702
                            },
                            {
                                "rc_agent": "fc1231d0-ccf1-42a9-b575-94ca5fa84d7b",
                                "total_count": 1112,
                                "total_size": 292564709
                            },
                            {
                                "rc_agent": "5928a144-0ec8-4f97-a120-f2d841092142",
                                "total_count": 725,
                                "total_size": 62509521
                            },
                            {
                                "rc_agent": "9b85f21e-6ab5-4d23-89a6-0040c1ec7752",
                                "total_count": 19031,
                                "total_size": 1787419062
                            },
                            {
                                "rc_agent": "69eb7810-490d-4bca-9d63-ce8ef09cdd73",
                                "total_count": 1120,
                                "total_size": 146681669
                            },
                            {
                                "rc_agent": "91542295-6a78-41e7-ab72-26d1151213eb",
                                "total_count": 12088,
                                "total_size": 1168488896
                            },
                            {
                                "rc_agent": "a5f951d6-dc38-4ce5-befe-3b9566913c8d",
                                "total_count": 1134,
                                "total_size": 122271572
                            },
                            {
                                "rc_agent": "371b17b9-a20b-4ece-87bf-217f3272c0bc",
                                "total_count": 2058,
                                "total_size": 527577833
                            },
                            {
                                "rc_agent": "43aea9ae-efc6-45f5-960a-00ae862d9164",
                                "total_count": 816,
                                "total_size": 73690743
                            },
                            {
                                "rc_agent": "18c58d7c-a65a-47f5-97e1-61be41faa47e",
                                "total_count": 522,
                                "total_size": 163817123
                            },
                            {
                                "rc_agent": "6e81d920-b542-4fb3-a5fe-a7b62ba775a5",
                                "total_count": 558,
                                "total_size": 263956857
                            },
                            {
                                "rc_agent": "fd7e7369-503c-470f-9935-9f0a5e66eab8",
                                "total_count": 59,
                                "total_size": 1632128
                            },
                            {
                                "rc_agent": "daf4cc81-cf66-4995-b35f-7fe1494b2d35",
                                "total_count": 1508,
                                "total_size": 134636663
                            },
                            {
                                "rc_agent": "add4a55a-a648-4e58-bbe4-3a30959efbf5",
                                "total_count": 195,
                                "total_size": 6431602
                            },
                            {
                                "rc_agent": "b2b5e0c2-3313-421c-8c24-24cdb9c9a84f",
                                "total_count": 95,
                                "total_size": 728680574
                            },
                            {
                                "rc_agent": "7ffc4d3f-9ee9-469a-9510-ea462d2c5e1a",
                                "total_count": 4436,
                                "total_size": 5995951012
                            },
                            {
                                "rc_agent": "bc9086c2-4a14-4b67-b2e5-2fded3f52c0d",
                                "total_count": 3560,
                                "total_size": 2515693320
                            },
                            {
                                "rc_agent": "a596e01c-e3c1-4325-8d45-8acc58865fba",
                                "total_count": 1831,
                                "total_size": 187312536
                            },
                            {
                                "rc_agent": "d2c65ba3-c2c6-47b5-b8c5-b3db6739aae5",
                                "total_count": 12885,
                                "total_size": 799742481
                            },
                            {
                                "rc_agent": "2b3ad81b-be5c-46f6-899f-95a5497d3451",
                                "total_count": 1163,
                                "total_size": 391020358
                            },
                            {
                                "rc_agent": "dd96a8b6-73f7-4a4a-90ff-5bddcce48668",
                                "total_count": 3268,
                                "total_size": 2341559015
                            },
                            {
                                "rc_agent": "7f089cb1-b47b-46ba-be1b-3ea0e4d34c30",
                                "total_count": 1808,
                                "total_size": 463659228
                            },
                            {
                                "rc_agent": "4632ffd5-f129-4521-9d46-14f7aec7ba5c",
                                "total_count": 1472,
                                "total_size": 349542285
                            },
                            {
                                "rc_agent": "b6b9cf88-bcfd-4ab5-9384-4062557d224f",
                                "total_count": 1008,
                                "total_size": 301351891
                            },
                            {
                                "rc_agent": "add53d85-dfc2-45ad-b82f-226a50e11302",
                                "total_count": 13414,
                                "total_size": 1188867629
                            },
                            {
                                "rc_agent": "5d7189fb-d9e5-4a01-bc87-e0b6bc77c5b5",
                                "total_count": 870,
                                "total_size": 451929689
                            },
                            {
                                "rc_agent": "09fc7f75-6bea-479c-aa5c-d48c64d20b00",
                                "total_count": 2114,
                                "total_size": 557541053
                            },
                            {
                                "rc_agent": "12445e91-de6f-408d-b418-a82710400d24",
                                "total_count": 1136,
                                "total_size": 441968465
                            },
                            {
                                "rc_agent": "96f1c46e-102f-4c85-9e61-4186493c9436",
                                "total_count": 314,
                                "total_size": 14250720
                            },
                            {
                                "rc_agent": "240970e3-1c02-43ea-bce7-dbac5ba2936c",
                                "total_count": 15,
                                "total_size": 1682784
                            },
                            {
                                "rc_agent": "76c766ae-2d4e-4782-9786-b9151f1fbd6b",
                                "total_count": 36,
                                "total_size": 1410328
                            },
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 2748,
                                "total_size": 1298144229
                            },
                            {
                                "rc_agent": "253b47a4-3926-479d-8c74-1283cb537806",
                                "total_count": 906,
                                "total_size": 26072969
                            },
                            {
                                "rc_agent": "5ea80738-f396-4573-ad58-316d367767e1",
                                "total_count": 1115,
                                "total_size": 368578166
                            },
                            {
                                "rc_agent": "29066717-fd7f-4192-b2e1-c0eb81472bed",
                                "total_count": 873,
                                "total_size": 223504600
                            },
                            {
                                "rc_agent": "48c704af-1d2c-4e28-815d-149c96d58746",
                                "total_count": 116,
                                "total_size": 70754830
                            },
                            {
                                "rc_agent": "f16be751-5e1e-4441-a1a4-e9a2265c95e3",
                                "total_count": 1037,
                                "total_size": 256409888
                            },
                            {
                                "rc_agent": "7cda642a-581c-423e-8a80-fe3bb0b395aa",
                                "total_count": 21209,
                                "total_size": 2020696068
                            },
                            {
                                "rc_agent": "cc15f98b-cf4f-46b7-ba0f-12655031b9f9",
                                "total_count": 1032,
                                "total_size": 77906980
                            },
                            {
                                "rc_agent": "43da00b4-4ab7-43a8-83ae-4256e77773e9",
                                "total_count": 956,
                                "total_size": 224031278
                            },
                            {
                                "rc_agent": "efb88be7-6b47-43b1-a5cb-5ae251164f32",
                                "total_count": 3086,
                                "total_size": 1745818999
                            },
                            {
                                "rc_agent": "5641f5e7-4ce2-4eb5-8fe7-84d5603e57c2",
                                "total_count": 3027,
                                "total_size": 2835873869
                            },
                            {
                                "rc_agent": "e4e38f4f-5c20-43e8-be7f-a2584ee039af",
                                "total_count": 18,
                                "total_size": 574150640
                            },
                            {
                                "rc_agent": "aee56be4-88c0-4c7a-b8a3-9f38324b1a54",
                                "total_count": 1674,
                                "total_size": 78151654
                            },
                            {
                                "rc_agent": "d7252561-0f3f-4ff8-962c-5aa28c348e88",
                                "total_count": 1310,
                                "total_size": 5329139371
                            },
                            {
                                "rc_agent": "ccfa643a-f98b-4b0a-acf8-e103b3cc4228",
                                "total_count": 1450,
                                "total_size": 73295289
                            },
                            {
                                "rc_agent": "3d53bde8-f658-40e0-9d5c-485efdd28b24",
                                "total_count": 36,
                                "total_size": 6542755
                            },
                            {
                                "rc_agent": "2e358d34-dac2-4275-b3b2-337a901fba29",
                                "total_count": 2340,
                                "total_size": 157104824
                            },
                            {
                                "rc_agent": "d7bee473-4b45-4209-9088-c98632e4e5b3",
                                "total_count": 1038,
                                "total_size": 358876573
                            },
                            {
                                "rc_agent": "5ab9becc-c5d5-4a0b-939e-46fa600dbdc4",
                                "total_count": 28829,
                                "total_size": 2615226693
                            },
                            {
                                "rc_agent": "1bb77442-1ede-4256-ac9d-295bb86dd0ab",
                                "total_count": 1231,
                                "total_size": 415507531
                            },
                            {
                                "rc_agent": "0c2916bf-ea16-4fab-9a02-9ede36521332",
                                "total_count": 3,
                                "total_size": 1533041
                            },
                            {
                                "rc_agent": "b93bc53f-5574-4fa6-9544-925cb968aa98",
                                "total_count": 6,
                                "total_size": 1636175
                            },
                            {
                                "rc_agent": "43ce5e2d-2113-4010-901c-c7be169d00b3",
                                "total_count": 3965,
                                "total_size": 2980580197
                            },
                            {
                                "rc_agent": "95c3d2f2-eff2-4c49-abd2-07f535701e11",
                                "total_count": 16416,
                                "total_size": 1491008693
                            },
                            {
                                "rc_agent": "c6ec97b3-9b87-4b16-a133-e7c3382ff7c3",
                                "total_count": 1071,
                                "total_size": 236829504
                            },
                            {
                                "rc_agent": "e66c3af0-c745-4a42-8d1e-756e9566d325",
                                "total_count": 1158,
                                "total_size": 1688445935
                            },
                            {
                                "rc_agent": "331ed874-758b-47a3-ab10-0959939a725d",
                                "total_count": 966,
                                "total_size": 37491491
                            },
                            {
                                "rc_agent": "d2c1ce33-9b5f-4428-9908-150f59e0e36e",
                                "total_count": 1360,
                                "total_size": 96604311
                            },
                            {
                                "rc_agent": "34a77975-c47d-49ac-968c-ad54d388a854",
                                "total_count": 868,
                                "total_size": 5425310066
                            },
                            {
                                "rc_agent": "7ce97de2-fdf8-4320-8170-e70edd111c38",
                                "total_count": 64,
                                "total_size": 14587025
                            },
                            {
                                "rc_agent": "54d9abb3-7272-4572-aec1-ead93c75b92e",
                                "total_count": 3933,
                                "total_size": 317264722
                            },
                            {
                                "rc_agent": "68b0044d-cc55-4fb0-8361-170949cf25a9",
                                "total_count": 293,
                                "total_size": 280971992
                            },
                            {
                                "rc_agent": "1acb96a2-3872-4512-8a76-b7c73ed85873",
                                "total_count": 648,
                                "total_size": 24224583
                            },
                            {
                                "rc_agent": "edd31932-a1a1-485a-809a-91bcc5125859",
                                "total_count": 89,
                                "total_size": 966858
                            },
                            {
                                "rc_agent": "b6f68c6d-be4f-423c-9497-f1fc3ec62285",
                                "total_count": 1149,
                                "total_size": 376081538
                            },
                            {
                                "rc_agent": "c9dbe159-0b53-480b-aedf-f9d07d21820f",
                                "total_count": 1112,
                                "total_size": 1104208123
                            },
                            {
                                "rc_agent": "e6dfd568-d18f-4e7a-aeb9-1a545c5d5d8e",
                                "total_count": 38,
                                "total_size": 22076614
                            },
                            {
                                "rc_agent": "62751ae0-b406-4283-9ade-0c2547853cdd",
                                "total_count": 8297,
                                "total_size": 673122914
                            },
                            {
                                "rc_agent": "603ba6de-95b7-401d-8ece-3abf62bab484",
                                "total_count": 1196,
                                "total_size": 306102819
                            },
                            {
                                "rc_agent": "43529067-ea39-4fe7-bc3b-9a4253360ee3",
                                "total_count": 824,
                                "total_size": 311645317
                            },
                            {
                                "rc_agent": "70d9d391-b2e7-4d1c-9d80-5d2fe43537c1",
                                "total_count": 32,
                                "total_size": 4090482
                            },
                            {
                                "rc_agent": "99411921-fb0a-4d98-9dc4-fa4a7b01be56",
                                "total_count": 168,
                                "total_size": 3104816
                            },
                            {
                                "rc_agent": "f669ddcd-3dca-4214-9326-95ab664fb4e8",
                                "total_count": 96,
                                "total_size": 4380630
                            },
                            {
                                "rc_agent": "8ee58126-0e28-4a09-bd09-8a86eedd8333",
                                "total_count": 100,
                                "total_size": 31316221
                            },
                            {
                                "rc_agent": "62d71b76-9384-4273-84e0-234cf5e51789",
                                "total_count": 9283,
                                "total_size": 726999119
                            },
                            {
                                "rc_agent": "4bb2b074-5e44-4741-b16a-925823525dc6",
                                "total_count": 1073,
                                "total_size": 417070903
                            },
                            {
                                "rc_agent": "1813a472-72bf-4994-bacc-078fadef2ef4",
                                "total_count": 4162,
                                "total_size": 2424833922
                            },
                            {
                                "rc_agent": "07ac0c92-3c6e-4069-a6cd-e52889c8ebfc",
                                "total_count": 528,
                                "total_size": 227194573
                            },
                            {
                                "rc_agent": "993ce8d6-08f8-4cc5-8d19-e86a3f1e70b3",
                                "total_count": 347,
                                "total_size": 14063774
                            },
                            {
                                "rc_agent": "85e2a872-f592-437c-8ecd-e35639bff8ec",
                                "total_count": 28,
                                "total_size": 1575119
                            },
                            {
                                "rc_agent": "eb0f99de-b782-41d9-b1f7-af715f4db011",
                                "total_count": 9839,
                                "total_size": 1001015908
                            },
                            {
                                "rc_agent": "f02840b1-dc49-40c0-a058-fbec08f9b023",
                                "total_count": 1550,
                                "total_size": 384565065
                            },
                            {
                                "rc_agent": "64336a84-2bfa-4105-9401-22b0b4456088",
                                "total_count": 3494,
                                "total_size": 10400400587
                            },
                            {
                                "rc_agent": "e6c676de-2d43-400f-aa8f-11b67f0e6699",
                                "total_count": 1512,
                                "total_size": 1358314919
                            },
                            {
                                "rc_agent": "51e8bff9-ff13-4701-8992-e905e868650f",
                                "total_count": 156,
                                "total_size": 1896603126
                            },
                            {
                                "rc_agent": "e2a86058-a17f-4a14-ba45-b83f051deaed",
                                "total_count": 664,
                                "total_size": 141153039
                            },
                            {
                                "rc_agent": "e43ef4ad-e361-4b94-b521-daee811a74c8",
                                "total_count": 1390,
                                "total_size": 527655453
                            },
                            {
                                "rc_agent": "71df0525-9d20-466c-90be-5a5282f5ffb9",
                                "total_count": 1080,
                                "total_size": 326693439
                            },
                            {
                                "rc_agent": "5277d7ad-3292-4040-b689-91ff69309535",
                                "total_count": 14844,
                                "total_size": 1289857616
                            },
                            {
                                "rc_agent": "98135d8f-a0f6-4936-9f89-ccafbab00cb1",
                                "total_count": 1224,
                                "total_size": 96265340
                            },
                            {
                                "rc_agent": "ebb2380b-9900-47e6-a1d1-1a841da6abf6",
                                "total_count": 1398,
                                "total_size": 473728379
                            },
                            {
                                "rc_agent": "016e0a2a-f31d-40ef-9d5b-ee548d2caf8a",
                                "total_count": 2873,
                                "total_size": 1356908370
                            },
                            {
                                "rc_agent": "76d9cc1c-b208-4749-a7d8-c664ea263d68",
                                "total_count": 115,
                                "total_size": 9725040
                            },
                            {
                                "rc_agent": "0d3304de-ca7d-4f1b-b23f-b98c1dbd1c4e",
                                "total_count": 1136,
                                "total_size": 1084089674
                            },
                            {
                                "rc_agent": "66cbba7b-4aae-43a2-a71a-2e95be769c5b",
                                "total_count": 1059,
                                "total_size": 428331209
                            },
                            {
                                "rc_agent": "e641bd75-a339-4359-81f1-b978cf2c6f6e",
                                "total_count": 3366,
                                "total_size": 2274360456
                            },
                            {
                                "rc_agent": "034fbcfb-b7d3-43cf-8072-4c14393d0aca",
                                "total_count": 620,
                                "total_size": 480588539
                            },
                            {
                                "rc_agent": "0841895d-56b5-4998-af02-b665572ea744",
                                "total_count": 320,
                                "total_size": 6230139
                            },
                            {
                                "rc_agent": "2678b010-e682-4710-91a0-1892320d4253",
                                "total_count": 376,
                                "total_size": 53081241
                            },
                            {
                                "rc_agent": "e44814f9-0de2-4b8e-859c-95f47244c54b",
                                "total_count": 266,
                                "total_size": 45613494
                            },
                            {
                                "rc_agent": "fd4fee68-c95c-444c-b62a-cfb9f5cb3034",
                                "total_count": 48,
                                "total_size": 3137991
                            },
                            {
                                "rc_agent": "a982a122-78dc-4048-800e-bd9e530bea84",
                                "total_count": 1445,
                                "total_size": 130459233
                            },
                            {
                                "rc_agent": "2576df94-ae7d-47aa-9f7f-4433bf886a29",
                                "total_count": 1236,
                                "total_size": 267582301
                            },
                            {
                                "rc_agent": "f91e868f-5ed6-4665-8e19-45b20fff62fd",
                                "total_count": 3486,
                                "total_size": 6457136572
                            },
                            {
                                "rc_agent": "8e37837c-5ee4-474c-a492-32dd0028f9be",
                                "total_count": 734,
                                "total_size": 457391911
                            },
                            {
                                "rc_agent": "03320a05-40a5-425b-b566-529ff431b5a6",
                                "total_count": 3568,
                                "total_size": 2287571522
                            },
                            {
                                "rc_agent": "085546b6-46b0-4980-abd5-c5cf8b758228",
                                "total_count": 10271,
                                "total_size": 1056120997
                            },
                            {
                                "rc_agent": "56fa81fc-0e6b-4180-ae6f-5fa912ffef45",
                                "total_count": 8601,
                                "total_size": 790778705
                            },
                            {
                                "rc_agent": "c805ecb0-c4b2-4931-9234-67def8ecef69",
                                "total_count": 245,
                                "total_size": 147608938
                            },
                            {
                                "rc_agent": "0cbf7fbc-9405-4587-a3f7-3b9e9dad1251",
                                "total_count": 1456,
                                "total_size": 112414063
                            },
                            {
                                "rc_agent": "bdf31889-7185-4e50-99fd-dc90117a6430",
                                "total_count": 1860,
                                "total_size": 132751480
                            },
                            {
                                "rc_agent": "95130719-7458-415c-9b5e-aafbcd6df7ea",
                                "total_count": 148,
                                "total_size": 1218579678
                            },
                            {
                                "rc_agent": "cf90ebda-eb74-4a86-92ee-f99f252f7558",
                                "total_count": 634,
                                "total_size": 188605244
                            },
                            {
                                "rc_agent": "3057ad07-c4f3-46fe-b428-f76e633cf5ae",
                                "total_count": 1816,
                                "total_size": 4527268626
                            },
                            {
                                "rc_agent": "1abc8bab-6791-457e-b2ed-333614868cfc",
                                "total_count": 2520,
                                "total_size": 993152913
                            },
                            {
                                "rc_agent": "0143a74d-9e52-4478-aba0-f0853a0d9f02",
                                "total_count": 1131,
                                "total_size": 270118137
                            },
                            {
                                "rc_agent": "1ea8d513-7540-4db9-adfa-2eb1e0a03d7c",
                                "total_count": 1469,
                                "total_size": 442935154
                            },
                            {
                                "rc_agent": "d17f3140-fcd3-4960-a1cd-57034d4aaece",
                                "total_count": 1248,
                                "total_size": 344032010
                            },
                            {
                                "rc_agent": "d8e682eb-dce1-43c2-a5b6-2d8d12afb045",
                                "total_count": 1256,
                                "total_size": 975263898
                            },
                            {
                                "rc_agent": "0b9b8173-bbf8-4ed1-b9f5-16c96dca13b2",
                                "total_count": 43,
                                "total_size": 25386527
                            },
                            {
                                "rc_agent": "a8aeb97b-e736-40b2-aeb1-5af89a6b6ea2",
                                "total_count": 1319,
                                "total_size": 3738707110
                            },
                            {
                                "rc_agent": "17cca798-49c3-45ab-9a18-ea7e3eb0763c",
                                "total_count": 480,
                                "total_size": 397436801
                            },
                            {
                                "rc_agent": "5de4a678-ad67-473a-847a-e55d40742211",
                                "total_count": 1288,
                                "total_size": 17129091
                            },
                            {
                                "rc_agent": "d3594679-5fc5-4f6d-90e2-c42f1abefbc2",
                                "total_count": 1087,
                                "total_size": 548806959
                            },
                            {
                                "rc_agent": "e655d590-f1c8-4253-a1aa-a2f9fd1dba0c",
                                "total_count": 1923,
                                "total_size": 197376813
                            },
                            {
                                "rc_agent": "da517cae-bb01-445d-897e-6afd039b2ae5",
                                "total_count": 4472,
                                "total_size": 2247395190
                            },
                            {
                                "rc_agent": "5e6320c1-ef77-4fc9-8f60-40a104c0d1b1",
                                "total_count": 295,
                                "total_size": 5321232
                            },
                            {
                                "rc_agent": "23182adf-c01b-4fcf-be87-1c820dd3b222",
                                "total_count": 1232,
                                "total_size": 66485417
                            },
                            {
                                "rc_agent": "4c85e760-0b67-4bb2-8276-ac4ac64ca7dd",
                                "total_count": 2136,
                                "total_size": 870975645
                            },
                            {
                                "rc_agent": "aff601d0-2e32-49fb-b1ec-0b820ac8e47f",
                                "total_count": 60,
                                "total_size": 6155362
                            },
                            {
                                "rc_agent": "7ec34209-61c1-4229-bcd4-9cac2475ac0d",
                                "total_count": 32,
                                "total_size": 388819191
                            },
                            {
                                "rc_agent": "d86df6e7-0364-4cfc-aeee-e4c54b535b31",
                                "total_count": 966,
                                "total_size": 92112316
                            },
                            {
                                "rc_agent": "e20475a0-68d4-414c-b5c8-b884dc7196e0",
                                "total_count": 1187,
                                "total_size": 101905888
                            },
                            {
                                "rc_agent": "a80cda47-b44b-4f93-aacb-ba4170668710",
                                "total_count": 276,
                                "total_size": 216579750
                            },
                            {
                                "rc_agent": "e2d20f9b-3378-40b6-b4d1-c81e93abf559",
                                "total_count": 3191,
                                "total_size": 1093624831
                            },
                            {
                                "rc_agent": "1a138928-02a2-45d3-8f8a-809cd1439cb8",
                                "total_count": 1420,
                                "total_size": 120907885
                            },
                            {
                                "rc_agent": "181b89f0-08f3-4984-9eed-d0688f81c2e9",
                                "total_count": 103,
                                "total_size": 4472913
                            },
                            {
                                "rc_agent": "15b6f0c2-7c0a-4425-ad5d-83f43dd5afb9",
                                "total_count": 72,
                                "total_size": 24604183
                            },
                            {
                                "rc_agent": "d3bbaecd-db97-418c-b10b-f9a7e31c40f3",
                                "total_count": 1256,
                                "total_size": 1003679446
                            },
                            {
                                "rc_agent": "fce712f1-d2f7-45f9-bc8e-aad64c01d7e6",
                                "total_count": 186,
                                "total_size": 324267385
                            },
                            {
                                "rc_agent": "cc836486-f7a3-4d68-8aa4-8165b888cbc8",
                                "total_count": 1096,
                                "total_size": 110960003
                            },
                            {
                                "rc_agent": "75f6f6fe-5a24-4d70-805f-df04b5279819",
                                "total_count": 176,
                                "total_size": 3685118
                            },
                            {
                                "rc_agent": "0f272c32-849a-4c36-9670-6f78adc6929d",
                                "total_count": 698,
                                "total_size": 163390797
                            },
                            {
                                "rc_agent": "956da5e7-ea44-4db8-87c9-30d4c1670e8c",
                                "total_count": 12867,
                                "total_size": 1139510703
                            },
                            {
                                "rc_agent": "6e46edd9-2981-4071-91f3-55cc463cc377",
                                "total_count": 10246,
                                "total_size": 972057530
                            },
                            {
                                "rc_agent": "ab98e0f2-4bdb-48ff-b11c-e13a38d8049b",
                                "total_count": 6764,
                                "total_size": 447218591
                            },
                            {
                                "rc_agent": "8eddab92-6e24-4be4-a18f-2c76cb4fb59a",
                                "total_count": 310,
                                "total_size": 60019728
                            },
                            {
                                "rc_agent": "88b45b63-339d-4eba-93bb-edf6048eb517",
                                "total_count": 780,
                                "total_size": 58697795
                            },
                            {
                                "rc_agent": "ed919cd0-37dd-4bd9-89c5-98ad79c66b1a",
                                "total_count": 648,
                                "total_size": 405935179
                            },
                            {
                                "rc_agent": "e0c52a4e-4ad7-4759-9cb5-1e3fdeb92139",
                                "total_count": 2735,
                                "total_size": 1567019916
                            },
                            {
                                "rc_agent": "e036a237-5c68-447e-807f-e625d0ef5747",
                                "total_count": 20,
                                "total_size": 474543654
                            },
                            {
                                "rc_agent": "7a98d7fb-5aaf-4f4f-a3aa-a978479eec03",
                                "total_count": 108,
                                "total_size": 1301354
                            },
                            {
                                "rc_agent": "9299cf04-ddf3-4c8f-9b5f-0a21a54508f6",
                                "total_count": 736,
                                "total_size": 510973581
                            },
                            {
                                "rc_agent": "8fcf6d7a-a403-4902-ab05-8bc12e3e67f7",
                                "total_count": 9,
                                "total_size": 10569
                            },
                            {
                                "rc_agent": "6faee253-470c-4308-83b2-e325ba670ac5",
                                "total_count": 1832,
                                "total_size": 5538162027
                            },
                            {
                                "rc_agent": "cc2dd55e-fefa-4f2b-b3c2-389aa818372a",
                                "total_count": 14912,
                                "total_size": 1389096332
                            },
                            {
                                "rc_agent": "2590ec90-bac7-4b82-a744-7be84a9bb344",
                                "total_count": 864,
                                "total_size": 5397592953
                            },
                            {
                                "rc_agent": "c22ac874-d374-4882-ab8f-b8dcf686a705",
                                "total_count": 1512,
                                "total_size": 1045945754
                            },
                            {
                                "rc_agent": "456d9bf9-a53a-4786-9c2c-24cdd5fd6771",
                                "total_count": 2423,
                                "total_size": 1075381010
                            },
                            {
                                "rc_agent": "9c6c550c-f498-4ca3-98b3-6eedbcde261e",
                                "total_count": 2321,
                                "total_size": 1172586580
                            },
                            {
                                "rc_agent": "24e602bd-4976-4971-ab39-f380f571cff0",
                                "total_count": 13775,
                                "total_size": 1335518151
                            },
                            {
                                "rc_agent": "cd694da2-6e2a-43e1-a815-071b4e6b4437",
                                "total_count": 62,
                                "total_size": 100629671
                            },
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 6871,
                                "total_size": 677334764
                            },
                            {
                                "rc_agent": "d71734a4-65e1-44b9-9edc-7c51b4829db1",
                                "total_count": 1080,
                                "total_size": 248628984
                            },
                            {
                                "rc_agent": "40fbe3e2-ecb9-4c4a-b7fe-83c01aa9f8d4",
                                "total_count": 10853,
                                "total_size": 860480858
                            },
                            {
                                "rc_agent": "8b29c58c-1389-43ca-b06c-ce5a4e22a592",
                                "total_count": 1937,
                                "total_size": 939669689
                            },
                            {
                                "rc_agent": "cce94ee9-b052-4958-b8ae-aae79fef3911",
                                "total_count": 849,
                                "total_size": 137731859
                            },
                            {
                                "rc_agent": "e0246fa1-f226-414d-ab79-dc401abc9163",
                                "total_count": 42,
                                "total_size": 720346
                            },
                            {
                                "rc_agent": "86a7146b-2784-4a4c-9400-c9f242bbc842",
                                "total_count": 120,
                                "total_size": 273752384
                            },
                            {
                                "rc_agent": "a90116fc-45a3-48bc-ab54-01b5ade652c7",
                                "total_count": 12,
                                "total_size": 1181886
                            },
                            {
                                "rc_agent": "58072d91-3e01-456f-b334-7774558af4e0",
                                "total_count": 1357,
                                "total_size": 222956220
                            },
                            {
                                "rc_agent": "11e76f00-f77a-4770-b102-494a7140fbdf",
                                "total_count": 64,
                                "total_size": 39822572
                            },
                            {
                                "rc_agent": "2c5bd94f-c19d-46c1-b9ae-e97f27d4e2a0",
                                "total_count": 972,
                                "total_size": 6063883424
                            },
                            {
                                "rc_agent": "57b4ffec-a5fa-459d-b74b-c6881325f6f9",
                                "total_count": 277,
                                "total_size": 27189059
                            },
                            {
                                "rc_agent": "099064fd-4727-46e0-8691-8e034d4807ed",
                                "total_count": 13075,
                                "total_size": 1111166519
                            },
                            {
                                "rc_agent": "bc5effd7-27c2-499e-975f-7f807ca5e1c0",
                                "total_count": 2006,
                                "total_size": 199131516
                            },
                            {
                                "rc_agent": "46d42c1f-5406-4852-9501-992cab5f2d4d",
                                "total_count": 4095,
                                "total_size": 871682077
                            },
                            {
                                "rc_agent": "b71bb2b7-e4cb-4edb-a4fb-5462220031ab",
                                "total_count": 16988,
                                "total_size": 1559689132
                            },
                            {
                                "rc_agent": "a4675878-6e3c-487d-bee4-4bb2c836be9d",
                                "total_count": 107,
                                "total_size": 1891651
                            },
                            {
                                "rc_agent": "77fccfef-0959-4714-9f3c-824ed28bee94",
                                "total_count": 551,
                                "total_size": 362553909
                            },
                            {
                                "rc_agent": "fc24910a-04d9-42c0-a210-815ae875bb71",
                                "total_count": 536,
                                "total_size": 293901828
                            },
                            {
                                "rc_agent": "69ea9348-0f58-42b2-be75-02ab7da491a6",
                                "total_count": 1928,
                                "total_size": 3228755373
                            },
                            {
                                "rc_agent": "fb4e78e6-5fb3-4793-b207-99b7fdc253aa",
                                "total_count": 855,
                                "total_size": 231740969
                            },
                            {
                                "rc_agent": "70e47b9a-8787-4d32-ba6f-b52bf4f3fda3",
                                "total_count": 1456,
                                "total_size": 109480934
                            },
                            {
                                "rc_agent": "df514785-8f0f-45c4-a387-47d788e01ff3",
                                "total_count": 96,
                                "total_size": 10964553
                            },
                            {
                                "rc_agent": "848e6c2c-a105-440d-abee-ecbe0c9606b5",
                                "total_count": 3448,
                                "total_size": 2194367622
                            },
                            {
                                "rc_agent": "6457d56c-7a7b-41bf-895e-2752b44eb8b0",
                                "total_count": 2189,
                                "total_size": 567637550
                            },
                            {
                                "rc_agent": "a4c416e7-1124-4e47-b696-22de16b51984",
                                "total_count": 480,
                                "total_size": 297200672
                            },
                            {
                                "rc_agent": "c174d8d5-099c-45d1-aa7c-c88d1328f33a",
                                "total_count": 1816,
                                "total_size": 5514605293
                            },
                            {
                                "rc_agent": "7077bf40-75d0-40dd-abcd-0802d13a86ce",
                                "total_count": 141,
                                "total_size": 1418934
                            },
                            {
                                "rc_agent": "d35bdf48-cf41-4422-acf4-176fe2c81da1",
                                "total_count": 873,
                                "total_size": 2976907302
                            },
                            {
                                "rc_agent": "3581f920-8bed-4d73-b069-04a03921b43b",
                                "total_count": 9,
                                "total_size": 35508605
                            },
                            {
                                "rc_agent": "22ea8678-ac0d-4c62-b3de-ffcffe2a7fa1",
                                "total_count": 9075,
                                "total_size": 806114443
                            },
                            {
                                "rc_agent": "b5fb7082-4934-40ea-baef-167226aadb2e",
                                "total_count": 1211,
                                "total_size": 76865120
                            },
                            {
                                "rc_agent": "7fe8e59a-b5c1-4e36-a9f4-b526023335fc",
                                "total_count": 21940,
                                "total_size": 2697904828
                            },
                            {
                                "rc_agent": "c4d9bb8f-fc98-47dd-9157-2175dd9a5a17",
                                "total_count": 1112,
                                "total_size": 728209102
                            },
                            {
                                "rc_agent": "b4c67ee2-7aef-4a6a-8b1b-58854f344d1f",
                                "total_count": 180,
                                "total_size": 10219110
                            },
                            {
                                "rc_agent": "81c02e77-1537-4338-a0c3-85db2d080d95",
                                "total_count": 8,
                                "total_size": 6846786
                            },
                            {
                                "rc_agent": "786e0af2-bb64-4108-8e6b-3dee9d6eaaea",
                                "total_count": 1115,
                                "total_size": 99581027
                            },
                            {
                                "rc_agent": "ce525686-da20-41f0-ab9c-a7251dfd6641",
                                "total_count": 1022,
                                "total_size": 409572274
                            },
                            {
                                "rc_agent": "8775d837-847a-4690-a4b6-8c0fa854ea62",
                                "total_count": 32,
                                "total_size": 1671062
                            },
                            {
                                "rc_agent": "7a239fb6-a845-471e-8c23-e39151363768",
                                "total_count": 10172,
                                "total_size": 864921088
                            },
                            {
                                "rc_agent": "8122c81a-f0dc-4440-82e9-b1c30b6c1acc",
                                "total_count": 12865,
                                "total_size": 1210847327
                            },
                            {
                                "rc_agent": "f37875b2-7d25-4b59-b7a6-72acf533c7f2",
                                "total_count": 463,
                                "total_size": 34168270
                            },
                            {
                                "rc_agent": "ad7cfef7-fb13-4b32-ab18-089dffcd367f",
                                "total_count": 1361,
                                "total_size": 37568184
                            },
                            {
                                "rc_agent": "65b0930f-bf2f-43f1-8f38-2fa8f9ae6d16",
                                "total_count": 1560,
                                "total_size": 1246697474
                            },
                            {
                                "rc_agent": "12df4716-7b9c-4f0c-ab9d-5f8c060513a1",
                                "total_count": 25162,
                                "total_size": 2573513153
                            },
                            {
                                "rc_agent": "8a80a7ab-9d46-4555-9d49-95769078b34c",
                                "total_count": 3891,
                                "total_size": 7855719688
                            },
                            {
                                "rc_agent": "8d1ab83e-60cb-4f11-907c-86719c4296b6",
                                "total_count": 1103,
                                "total_size": 236942587
                            },
                            {
                                "rc_agent": "dcfd7b1d-39de-490c-996c-6b7e1d657a5b",
                                "total_count": 2561,
                                "total_size": 1043740581
                            },
                            {
                                "rc_agent": "b3821a9e-6220-49ed-9676-a6646b114e15",
                                "total_count": 88,
                                "total_size": 67350645
                            },
                            {
                                "rc_agent": "739cdb10-526f-42a4-bd4e-3606184aed7f",
                                "total_count": 992,
                                "total_size": 69444683
                            },
                            {
                                "rc_agent": "25e8b1f8-3319-407c-84fd-9f04b2fa7c94",
                                "total_count": 909,
                                "total_size": 41954142
                            },
                            {
                                "rc_agent": "a4f9b16c-06c1-4b72-af20-6d2d051c0040",
                                "total_count": 1466,
                                "total_size": 57173621
                            },
                            {
                                "rc_agent": "29b38729-5ed9-4287-9460-af4a30be5773",
                                "total_count": 1329,
                                "total_size": 102953817
                            },
                            {
                                "rc_agent": "3bac79f3-7229-4293-8405-7c16556a3e98",
                                "total_count": 126,
                                "total_size": 252993761
                            },
                            {
                                "rc_agent": "b250cfa9-475d-4656-a3da-bd05c8064875",
                                "total_count": 88,
                                "total_size": 68972372
                            },
                            {
                                "rc_agent": "6a0bc126-5fef-494c-ace6-31e7ce209f77",
                                "total_count": 952,
                                "total_size": 129923526
                            },
                            {
                                "rc_agent": "9c07ff43-25bc-464b-8c5c-336eff9d67ca",
                                "total_count": 776,
                                "total_size": 95459765
                            }
                        ]
                    },
                    "time": "2025-02-20T02:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "1LySJJUBk0bmQQGuSRjp",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T17:00:00.000000Z",
                    "indices": {
                        "total_count": 8,
                        "total_size": 551835,
                        "group_by_agent": [
                            {
                                "rc_agent": "e77d580a-0219-4274-aa31-34be22fd05a8",
                                "total_count": 8,
                                "total_size": 551835
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "Ldw2JZUBWlvdGmPA9qZO",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T20:00:00.000000Z",
                    "indices": {
                        "total_count": 10,
                        "total_size": 3829915,
                        "group_by_agent": [
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 10,
                                "total_size": 3829915
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "3LykJZUBk0bmQQGusxiW",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T22:00:00.000000Z",
                    "indices": {
                        "total_count": 9,
                        "total_size": 4978623,
                        "group_by_agent": [
                            {
                                "rc_agent": "10183b80-363d-4a70-ad88-c4324346241c",
                                "total_count": 9,
                                "total_size": 4978623
                            }
                        ]
                    },
                    "time": "2025-02-20T11:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "XUYAJZUBT5H4H6HfCzau",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 215,
                        "total_size": 20647102,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 215,
                                "total_size": 20647102
                            }
                        ]
                    },
                    "time": "2025-02-20T07:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "ZEYAJZUBT5H4H6HfCzau",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 126,
                        "total_size": 12161920,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 126,
                                "total_size": 12161920
                            }
                        ]
                    },
                    "time": "2025-02-20T04:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "ZUYAJZUBT5H4H6HfCzau",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-20T19:00:00.000000Z",
                    "indices": {
                        "total_count": 289,
                        "total_size": 27785388,
                        "group_by_agent": [
                            {
                                "rc_agent": "e80c145d-c5c2-4286-a0dc-a1114839d860",
                                "total_count": 289,
                                "total_size": 27785388
                            }
                        ]
                    },
                    "time": "2025-02-20T03:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "ZtyAJpUBWlvdGmPArqa3",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-21T02:00:00.000000Z",
                    "indices": {
                        "total_count": 12,
                        "total_size": 1060161,
                        "group_by_agent": [
                            {
                                "rc_agent": "c18624e7-e7f5-48c1-9352-4965a0f8c59e",
                                "total_count": 12,
                                "total_size": 1060161
                            }
                        ]
                    },
                    "time": "2025-02-20T07:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            },
            {
                "_index": "ocpf-metadata-fixedreport-100000",
                "_type": "_doc",
                "_id": "aNyAJpUBWlvdGmPArqa3",
                "_score": 1.0039139,
                "_source": {
                    "belong_index": "ocpf-eos-metadata-100000-100003-43",
                    "belong_mtime": "2025-02-21T02:00:00.000000Z",
                    "indices": {
                        "total_count": 7,
                        "total_size": 602734,
                        "group_by_agent": [
                            {
                                "rc_agent": "c18624e7-e7f5-48c1-9352-4965a0f8c59e",
                                "total_count": 7,
                                "total_size": 602734
                            }
                        ]
                    },
                    "time": "2025-02-20T08:00:00.000000Z",
                    "time_field": "m_rcmtime"
                }
            }
        ]`

	data := make([]fixedreportData, 0)
	err := json.Unmarshal([]byte(dataStr), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	total := 0
	for _, v := range data {
		for _, a := range v.Source.Indices.GroupByAgent {
			if a.RcAgent == "eb195d39-20a2-4861-a995-214fc873f3a3" {
				total += a.TotalCount
			}
		}
	}
	fmt.Println(total)

}

type fixedreportData struct {
	Source *source `json:"_source"`
}

type source struct {
	BelongIndex string   `json:"belong_index"`
	BelongMtime string   `json:"belong_mtime"`
	Time        string   `json:"time"`
	TimeField   string   `json:"time_field"`
	Indices     *indices `json:"indices"`
}

type indices struct {
	GroupByAgent []*groupByAgent `json:"group_by_agent"`
}

type groupByAgent struct {
	RcAgent    string `json:"rc_agent"`
	TotalCount int    `json:"total_count"`
}
