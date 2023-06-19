# Various terms

Regarding email:

- **MX** : MX record or a mail exchange record is a DNS record that routes emails to the specified email servers. Essentially point to a IP address of the mail server's domain.
- **SPF** : Sender policy framework is a email authentication standard that helps protect users from spam/spoof/phishing attempts. Adding a SPF to a DNS allows users of that domain to send mail from that domain. Sort of like a permission to the users of that domain to send email.
- **DMARC** : Domain-based Message Authentication, Reporting & Conformance is the sort-of an opposite of an SPF. It protects users from emails. It handles emails either authenticated using a SPF or DKIM or not authenticated. 
 