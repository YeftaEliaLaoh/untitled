# untitled1
1.
A Soldier want magazine needed to be full ammo before war.
A Soldier will load bullets randomly into the magazine in gun.
If:
a) Soldier can put multiple magazine into one gun.
b) Only one magazine will be mark as verified, and soldier will stop put/remove magazine until it
verified.
c) Verified will happen when the magazine is full.
d) You must test the magazine is full or not with firing it on a gun.
How you can handle this in on programming?
Can you write the API based on those problem?

2.
Your mom has an online store name is "Kitara Store".
She has one big problem, in many time customer order requested, the products will produce minus
quantity.
When the product stock is limited and many customers order in the same time, the issue always occur.
Can you help to solve the problem on Kitara Store?
How you handle the race condition on Kitara Store?
Create API to your solution for proving your solution working.3.

3.
Joni menititipkan kunci rumahnya pada seorang temannya bernama Andi.
Namun, Andi memberikan sebuah games kepada Joni untuk mencari kunci rumahnya.
Joni ditandai dengan titik X.
Untuk menemukan kunci rumahnya, Andi memberikan beberapa petunjuk kepada Joni yaitu:
1. Jalan ke utara sebanyak Y langkah.
2. Lalu jalan ke timur sebanyak Y langkah.
3. Terkahir ke selatan sebanyak Y langkah.
Variable ‘Y’ menandakan angka yang hilang.
Buktikan berapa banyak titik yang menjadi kemungkinan lokasi kunci rumah Joni.
Jika diketahui angka yang hilang merupakan bilangan bulat positif.

#
1. /question1 This api handles random request bullets, and is equalize with random magazine data (0-5),
if request bullets equal with magazine the magazine is full
2. /question2 This api handles when the produce is limit , produce will sleep for 1 second and give 1 additional products 
3. /question3 This api to find the point that might be the key location of Joni's house

How to test
1. Run the code go run main.go
2. Execute the code
Example 
for question no 1
curl -X POST -H 'Content-Type: application/json' -i http://localhost:8080/question1 --data '{"Bullet": "1"}
{"Bullet": "2"}
{"Bullet": "3"}
{"Bullet": "4"}
{"Bullet": "5"}'

for question no 2
curl -X POST -H 'Content-Type: application/json' -i http://localhost:8080/question2 --data '{"Order": "1"}
{"Order": "2"}
'

for question no 3
curl -X POST -H 'Content-Type: application/json' -i http://localhost:8080/question3

