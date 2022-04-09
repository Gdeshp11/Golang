#!/bin/bash



list()
{
    echo "list of items in database: "
    curl -w "\n" "http://$IP:$PORT/list"
}

price()
{
    echo "price of item: $1"
    curl -w "\n" "http://$IP:$PORT/price?item=$1"
}

create()
{
    echo "adding item: $1 in database of price: $2"
    curl -w "\n" "http://$IP:$PORT/create?item=$1&price=$2"
}

update()
{
    echo "updating price of item: $1 with price: $2"
    curl -w "\n" "http://$IP:$PORT/update?item=$1&price=$2"
}

delete()
{
    echo "deleting item: $1 from database"
    curl -w "\n" "http://$IP:$PORT/delete?item=$1"
}

usage()
{
  echo "usage: test.sh -i <IP> -p <port>"
}


while getopts i:p: flag
do
    case "${flag}" in
        i) IP=${OPTARG};;
        p) PORT=${OPTARG};;
        *) exit 
    esac
done

if [ "$#" -eq 0 ]; then
  usage
  exit 1
fi

echo "IP: $IP";
echo "Port: $PORT";


printf "Enter input--->\nlist\nprice\ncreate\nupdate\ndelete\n"
read OPTION
case $OPTION in
  list)
    list
    ;;

  price)
    read -p "Enter item: " ITEM
    price $ITEM
    ;;

  create)
    read -p "Enter item: " ITEM
    echo "Enter price of item: "
    read PRICE
    create $ITEM $PRICE
    ;;

  update)
    read -p "Enter item: " ITEM
    echo "Enter price of item: "
    read PRICE
    update $ITEM $PRICE
    ;;

    delete)
    read -p "Enter item: " ITEM
    delete $ITEM
    ;;

  *)
    echo -n "unknown option: $OPTION"
    ;;
esac