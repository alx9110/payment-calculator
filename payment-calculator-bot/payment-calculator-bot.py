import telebot
import requests
import os


bot = telebot.TeleBot(os.getenv("TELEGRAM_BOT_TOKEN"), parse_mode=None) # You can set parse_mode by default. HTML or MARKDOWN

start = """
Payment Calculator by Alexander Telkov.

/records Просмотр показаний счетчиков
/taxes Просмотр тарифов по счетчикам
/new Ввод новых показаний
/update Обновление тарифов
"""

@bot.message_handler(commands=['start', 'help'], func=lambda message: True)
def send_welcome(message):
    print(message.from_user.id)
    bot.reply_to(message, start)

@bot.message_handler(commands=['records', 'help'])
def get_records(message):
    data = requests.get("http://localhost:80/api/records").json()
    text = "Просмотр показаний за последние полгода:\n"
    for line in data['data']:
        text += f" \
            Горячая вода: {line['HotValue']} {line['HotCost']}\n \
            Холодная вода: {line['ColdValue']} {line['ColdCost']}\n \
        "
    bot.reply_to(message, text)

@bot.message_handler(commands=['taxes', 'help'])
def get_records(message):
    data = requests.get("http://localhost:80/api/taxes").json()
    bot.reply_to(message, data)

@bot.message_handler(commands=['hot', 'help'], func=lambda message: True)
def create_records(message):
    print(message.from_user.id)
    rec = {
        "hot_price": 0.0,
        "cold_price": 2.5,
        "energy_price": 1.3,
        "drenage_price": 2.33,
    }
    headers = {"Content-Type": "application/json"}
    data = requests.post("http://localhost:80/api/taxes", json=rec, headers=headers).text
    print(requests.post("http://localhost:80/api/taxes", json=rec, headers=headers).request.body)
    bot.reply_to(message, data)

@bot.message_handler(func=lambda message: True)
def echo_all(message):
    bot.reply_to(message, message.text)

bot.infinity_polling()
