from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from plugin.models.base_model import Model
from plugin.models.exploit import Exploit
from plugin.models.info_finder import InfoFinder
from plugin.models.malware import Malware
from plugin.models.misconfiguration import Misconfiguration
from plugin.models.package import Package
from plugin.models.rootkit import Rootkit
from plugin.models.secret import Secret
from plugin.models.vulnerability import Vulnerability
from plugin import util

from plugin.models.exploit import Exploit  # noqa: E501
from plugin.models.info_finder import InfoFinder  # noqa: E501
from plugin.models.malware import Malware  # noqa: E501
from plugin.models.misconfiguration import Misconfiguration  # noqa: E501
from plugin.models.package import Package  # noqa: E501
from plugin.models.rootkit import Rootkit  # noqa: E501
from plugin.models.secret import Secret  # noqa: E501
from plugin.models.vulnerability import Vulnerability  # noqa: E501

class VMClarityData(Model):
    """NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).

    Do not edit the class manually.
    """

    def __init__(self, packages=None, vulnerabilities=None, malware=None, secrets=None, misconfigurations=None, rootkits=None, exploits=None, info_finder=None):  # noqa: E501
        """VMClarityData - a model defined in OpenAPI

        :param packages: The packages of this VMClarityData.  # noqa: E501
        :type packages: List[Package]
        :param vulnerabilities: The vulnerabilities of this VMClarityData.  # noqa: E501
        :type vulnerabilities: List[Vulnerability]
        :param malware: The malware of this VMClarityData.  # noqa: E501
        :type malware: List[Malware]
        :param secrets: The secrets of this VMClarityData.  # noqa: E501
        :type secrets: List[Secret]
        :param misconfigurations: The misconfigurations of this VMClarityData.  # noqa: E501
        :type misconfigurations: List[Misconfiguration]
        :param rootkits: The rootkits of this VMClarityData.  # noqa: E501
        :type rootkits: List[Rootkit]
        :param exploits: The exploits of this VMClarityData.  # noqa: E501
        :type exploits: List[Exploit]
        :param info_finder: The info_finder of this VMClarityData.  # noqa: E501
        :type info_finder: List[InfoFinder]
        """
        self.openapi_types = {
            'packages': List[Package],
            'vulnerabilities': List[Vulnerability],
            'malware': List[Malware],
            'secrets': List[Secret],
            'misconfigurations': List[Misconfiguration],
            'rootkits': List[Rootkit],
            'exploits': List[Exploit],
            'info_finder': List[InfoFinder]
        }

        self.attribute_map = {
            'packages': 'packages',
            'vulnerabilities': 'vulnerabilities',
            'malware': 'malware',
            'secrets': 'secrets',
            'misconfigurations': 'misconfigurations',
            'rootkits': 'rootkits',
            'exploits': 'exploits',
            'info_finder': 'infoFinder'
        }

        self._packages = packages
        self._vulnerabilities = vulnerabilities
        self._malware = malware
        self._secrets = secrets
        self._misconfigurations = misconfigurations
        self._rootkits = rootkits
        self._exploits = exploits
        self._info_finder = info_finder

    @classmethod
    def from_dict(cls, dikt) -> 'VMClarityData':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The VMClarityData of this VMClarityData.  # noqa: E501
        :rtype: VMClarityData
        """
        return util.deserialize_model(dikt, cls)

    @property
    def packages(self) -> List[Package]:
        """Gets the packages of this VMClarityData.


        :return: The packages of this VMClarityData.
        :rtype: List[Package]
        """
        return self._packages

    @packages.setter
    def packages(self, packages: List[Package]):
        """Sets the packages of this VMClarityData.


        :param packages: The packages of this VMClarityData.
        :type packages: List[Package]
        """

        self._packages = packages

    @property
    def vulnerabilities(self) -> List[Vulnerability]:
        """Gets the vulnerabilities of this VMClarityData.


        :return: The vulnerabilities of this VMClarityData.
        :rtype: List[Vulnerability]
        """
        return self._vulnerabilities

    @vulnerabilities.setter
    def vulnerabilities(self, vulnerabilities: List[Vulnerability]):
        """Sets the vulnerabilities of this VMClarityData.


        :param vulnerabilities: The vulnerabilities of this VMClarityData.
        :type vulnerabilities: List[Vulnerability]
        """

        self._vulnerabilities = vulnerabilities

    @property
    def malware(self) -> List[Malware]:
        """Gets the malware of this VMClarityData.


        :return: The malware of this VMClarityData.
        :rtype: List[Malware]
        """
        return self._malware

    @malware.setter
    def malware(self, malware: List[Malware]):
        """Sets the malware of this VMClarityData.


        :param malware: The malware of this VMClarityData.
        :type malware: List[Malware]
        """

        self._malware = malware

    @property
    def secrets(self) -> List[Secret]:
        """Gets the secrets of this VMClarityData.


        :return: The secrets of this VMClarityData.
        :rtype: List[Secret]
        """
        return self._secrets

    @secrets.setter
    def secrets(self, secrets: List[Secret]):
        """Sets the secrets of this VMClarityData.


        :param secrets: The secrets of this VMClarityData.
        :type secrets: List[Secret]
        """

        self._secrets = secrets

    @property
    def misconfigurations(self) -> List[Misconfiguration]:
        """Gets the misconfigurations of this VMClarityData.


        :return: The misconfigurations of this VMClarityData.
        :rtype: List[Misconfiguration]
        """
        return self._misconfigurations

    @misconfigurations.setter
    def misconfigurations(self, misconfigurations: List[Misconfiguration]):
        """Sets the misconfigurations of this VMClarityData.


        :param misconfigurations: The misconfigurations of this VMClarityData.
        :type misconfigurations: List[Misconfiguration]
        """

        self._misconfigurations = misconfigurations

    @property
    def rootkits(self) -> List[Rootkit]:
        """Gets the rootkits of this VMClarityData.


        :return: The rootkits of this VMClarityData.
        :rtype: List[Rootkit]
        """
        return self._rootkits

    @rootkits.setter
    def rootkits(self, rootkits: List[Rootkit]):
        """Sets the rootkits of this VMClarityData.


        :param rootkits: The rootkits of this VMClarityData.
        :type rootkits: List[Rootkit]
        """

        self._rootkits = rootkits

    @property
    def exploits(self) -> List[Exploit]:
        """Gets the exploits of this VMClarityData.


        :return: The exploits of this VMClarityData.
        :rtype: List[Exploit]
        """
        return self._exploits

    @exploits.setter
    def exploits(self, exploits: List[Exploit]):
        """Sets the exploits of this VMClarityData.


        :param exploits: The exploits of this VMClarityData.
        :type exploits: List[Exploit]
        """

        self._exploits = exploits

    @property
    def info_finder(self) -> List[InfoFinder]:
        """Gets the info_finder of this VMClarityData.


        :return: The info_finder of this VMClarityData.
        :rtype: List[InfoFinder]
        """
        return self._info_finder

    @info_finder.setter
    def info_finder(self, info_finder: List[InfoFinder]):
        """Sets the info_finder of this VMClarityData.


        :param info_finder: The info_finder of this VMClarityData.
        :type info_finder: List[InfoFinder]
        """

        self._info_finder = info_finder
