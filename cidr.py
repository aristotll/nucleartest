import netaddr

new_cidr = "10.233.64.0/18"
old_cidr = "10.233.0.0/18"

def is_overlap(old, new):
    old_net = netaddr.IPNetwork(old)
    new_net = netaddr.IPNetwork(new)
    return old_net in new_net or new_net in old_net

def __main__():
    print(is_overlap(old_cidr, new_cidr))

if __name__ == "__main__":
    __main__()
