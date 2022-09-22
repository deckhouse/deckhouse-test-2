#!/bin/bash

log_file="$1"
comment_url="$2"

if [ -z "$log_file" ]; then
  echo "Log file is required"
  exit 1
fi

if [ -z "$comment_url" ]; then
  echo "Comment url is required"
  exit 1
fi

master_ip=""
master_user=""

# wait master ip and user. 10 minutes 60 cycles wit 10 second sleep
for (( i=1; i<=60; i++ )); do
  # yep sleep before
  sleep 10

  ip=""
  if ! ip="$(grep -Po '(?<=master_ip_address_for_ssh = ).+$' "$log_file")"; then
    echo "Master ip not found"
    continue
  fi

  #https://stackoverflow.com/posts/36760050/revisions
  # we need to verify ip because string ca fsynced partially
  if ! echo "$ip" | grep -Po '((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}'; then
    echo "$ip is not ip"
    continue
  fi

  master_ip=$ip
  echo "IP found $master_ip"

  user=""
  if ! user="$(grep -Po '(?<=master_user_name_for_ssh = ).+$' "$log_file")"; then
    echo "User not found"
    continue
  fi

  if [ -z "$user" ]; then
    continue
  fi

  master_user="$user"
  echo "User was found: $master_user"

  # got ip and user
  break
done

if [[ "$master_ip" == "" || "$master_user" == "" ]]; then
  echo "Timeout waiting master ip and master user"
  exit 1
fi

body=""
# get body
for (( i=1; i<=5; i++ )); do
  got_body="$(curl \
    -f \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer $GITHUB_TOKEN" \
    "$comment_url"
  )"
  exit_code="$?"
  if [ "$exit_code" == 0 ]; then
    echo "Comment result $body"

    if ! bbody="$(echo "$got_body" | jq -r '.body')"; then
        continue
    fi

    body="${bbody}\n\nMaster ssh connection string: ssh ${master_user}@${master_ip}\n"
    break
  fi

  sleep 5
done

if [ -z "$body" ]; then
  echo "Timeout waiting comment body"
  exit 1
fi


# update comment
for (( i=1; i<=5; i++ )); do
  response_file=$(mktemp)
  http_code="$(curl \
    --output "$response_file" \
    --write-out "%{http_code}" \
    -X PATCH \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer $GITHUB_TOKEN" \
    -d "{\"body\":\"$body\"}" \
    "$comment_url"
  )"
  exit_code="$?"

  cat "$response_file"
  rm -f "$response_file"

  if [ "$exit_code" == 0 ]; then
    if [ "$http_code" == "200" ]; then
        exit 0
    fi

    echo "Comment not updated, http code: $http_code"
  fi

  sleep 5
done

echo "Timeout waiting comment updating"
exit 1
